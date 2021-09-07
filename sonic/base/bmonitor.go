package base

import (
	"context"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/fasttls"
	monitor "github.com/ProjectAthenaa/sonic-core/protos/monitorController"
	proxy_rater "github.com/ProjectAthenaa/sonic-core/protos/proxy-rater"
	"github.com/ProjectAthenaa/sonic-core/sonic/core"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/product"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	"github.com/go-redis/redis/v8"
	"github.com/json-iterator/go"
	"github.com/prometheus/common/log"
	"github.com/viney-shih/go-lock"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

var (
	json         = jsoniter.ConfigCompatibleWithStandardLibrary
	monitorCount = os.Getenv("MONITOR_TASK_COUNT")
)

type BMonitor struct {
	Data     *monitor.Task
	Ctx      context.Context
	Callback face.MonitorCallback
	Client   *fasttls.Client
	Monitor  struct {
		Channel chan map[string]interface{}
	}

	cancel context.CancelFunc

	redisKey      string
	proxyRedisKey string
	site          string
	rdb           redis.UniversalClient
	proxy         proxy
	proxyClient   proxy_rater.ProxyRaterClient

	//prv
	_proxyLocker lock.Mutex //mutex lock to avoid mismatches between authorization proxy and transport proxy
}

//used for proxies
type proxy struct {
	address    string
	authHeader string
}

func (tk *BMonitor) Listen() {
	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-c:
			tk.Stop()
			return
		case <-tk.Ctx.Done():
			return
		default:
			count := core.Base.GetRedis("cache").PubSubNumSub(tk.Ctx, tk.Data.RedisChannel).Val()
			if v, ok := count[tk.Data.RedisChannel]; v == 0 || !ok {
				tk.Stop()
			}
			time.Sleep(time.Second)
		}
	}
}

func (tk *BMonitor) Start(site product.Site, client proxy_rater.ProxyRaterClient) error {
	tk.Ctx, tk.cancel = context.WithCancel(context.Background())

	if tk.Data == nil {
		return face.ErrTaskHasNoData
	}

	taskCount, err := strconv.Atoi(monitorCount)
	if err != nil {
		return err
	}

	tk.Callback.OnStarting()

	if tk.Client == nil {
		tk.Client = fasttls.DefaultClient
	}

	tk.rdb = core.Base.GetRedis("cache")
	tk.redisKey = fmt.Sprintf(tk.Data.RedisChannel)
	tk.proxyRedisKey = fmt.Sprintf("proxies:monitors:%s", tk.Data.Site)
	tk._proxyLocker = lock.NewCASMutex()
	tk.proxyClient = client
	tk.site = string(site)
	tk.Monitor.Channel = make(chan map[string]interface{})

	var proxyWait sync.WaitGroup
	proxyWait.Add(1)
	go tk.proxyRefresher(&proxyWait)
	proxyWait.Wait()

	for i := 0; i < taskCount; i++ {
		go tk.Callback.TaskLoop()
	}

	go tk.Listen()
	go tk.submit()
	return nil
}

func (tk *BMonitor) Stop() {
	tk.cancel()
	tk.Callback.OnStopping()
}

func (tk *BMonitor) submit() {
	for {
		select {
		case msg := <-tk.Monitor.Channel:
			go func() {
				payload, err := json.Marshal(&msg)
				if err != nil {
					log.Error("error serialising data", err)
				}

				log.Info(string(payload))
				tk.rdb.Publish(tk.Ctx, tk.redisKey, string(payload))
			}()
		case <-tk.Ctx.Done():
			return
		default:
			continue
		}
	}
}

func (tk *BMonitor) proxyRefresher(wg *sync.WaitGroup) {
	var firstCalculated bool

	for range time.Tick(time.Second) {
		select {
		case <-tk.Ctx.Done():
			return
		default:
			break
		}
		tk._proxyLocker.Lock()
		proxyResp, err := tk.proxyClient.GetProxy(tk.Ctx, &proxy_rater.Site{Value: tk.site})
		if err != nil {
			log.Error("get proxy req: ", err)
			goto onErrorContinue
		}

		tk.proxy = proxy{
			address:    proxyResp.Value,
			authHeader: proxyResp.Authorization,
		}

		if !firstCalculated {
			wg.Done()
			firstCalculated = true
		}

	onErrorContinue:
		tk._proxyLocker.Unlock()
		continue
	}

}

func (tk *BMonitor) NewRequest(method, url string, body []byte) (*fasttls.Request, error) {
	return tk.Client.NewRequest(fasttls.Method(strings.ToUpper(method)), url, body)
}

func (tk *BMonitor) Do(req *fasttls.Request) (*fasttls.Response, error) {
	return tk.Client.DoCtx(tk.Ctx, req)
}
