package base

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/protos/monitor"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	"github.com/go-redis/redis/v8"
	"github.com/json-iterator/go"
	"github.com/prometheus/common/log"
	"github.com/viney-shih/go-lock"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)


var (
	json         = jsoniter.ConfigCompatibleWithStandardLibrary
	monitorCount = os.Getenv("MONITOR_TASK_COUNT")
)

const (
	posInf = "+inf"
	negInf = "-inf"
)

type BMonitor struct {
	Data     *monitor.Task
	Ctx      context.Context
	Callback face.MonitorCallback
	Client   *http.Client

	cancel context.CancelFunc

	redisKey      string
	proxyRedisKey string
	site          string
	rdb           *redis.Client
	proxy         proxy

	//prv
	_proxyLocker lock.Mutex //mutex lock to avoid mismatches between authorization proxy and transport proxy
}

//used for proxies
type proxy struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
	IP       string  `json:"ip"`
	Port     string  `json:"port"`
}

func (tk *BMonitor) Listen() {
	pubSub := tk.rdb.Subscribe(tk.Ctx, tk.redisKey)
	channel := pubSub.Channel()
	defer func() {
		if err := pubSub.Close(); err != nil {
			log.Error("error closing pubsub", err)
		}
	}()

	for {
		select {
		case <-tk.Ctx.Done():
			log.Info("monitor context timeout")
			return
		case msg := <-channel:
			if msg.Payload == "STOP" {
				log.Info("stop command received, stopping..")
				tk.Stop()
			}
		default:
			continue
		}
	}
}

func (tk *BMonitor) Start() error {
	if tk.Data == nil {
		return face.ErrTaskHasNoData
	}

	taskCount, err := strconv.Atoi(monitorCount)
	if err != nil {
		return err
	}

	tk.Callback.OnStarting()

	if tk.Client == nil {
		tk.Client = http.DefaultClient
	}

	tk.redisKey = fmt.Sprintf("monitors:%s", tk.Data.RedisChannelName)
	tk.proxyRedisKey = fmt.Sprintf("proxies:monitors:%s", tk.Data.Site)
	tk._proxyLocker = lock.NewCASMutex()

	if tk.cancel == nil {
		tk.Ctx, tk.cancel = context.WithCancel(tk.Ctx)
	}

	var proxyWait sync.WaitGroup
	proxyWait.Add(1)
	go tk.proxyRefresher(&proxyWait)
	proxyWait.Wait()

	for i := 0; i < taskCount; i++ {
		go tk.Callback.TaskLoop()
	}

	go tk.Listen()
	return nil
}

func (tk *BMonitor) Stop() {
	tk.cancel()
	tk.Callback.OnStopping()
}

func (tk *BMonitor) Submit(data map[string]interface{}) error {
	payload, err := json.Marshal(&data)
	if err != nil {
		log.Error("error serialising data", err)
		return err
	}

	tk.rdb.Publish(tk.Ctx, tk.redisKey, string(payload))
	return nil
}

func (tk *BMonitor) proxyRefresher(wg *sync.WaitGroup) {
	var currentIndex int
	var maxIndex = tk.rdb.ZCount(tk.Ctx, tk.proxyRedisKey, negInf, posInf).Val()

	var pr proxy
	var proxyUrl *url.URL

	var firstCalculated bool

	for range time.Tick(time.Second) {
		currentIndex++
		if currentIndex > int(maxIndex) {
			currentIndex = 1
		}

		proxyData, err := tk.rdb.ZRange(tk.Ctx, tk.proxyRedisKey, maxIndex-1, maxIndex).Result()
		if err != nil {
			log.Error("error completing zrange func", err)
			continue
		}

		tk._proxyLocker.TryLockWithContext(tk.Ctx)
		if err = json.Unmarshal([]byte(proxyData[0]), &pr); err != nil {
			log.Error("error unmarshaling proxy data", err)
			goto onErrorContinue
		}

		if pr.Username != nil && pr.Password != nil {
			proxyUrl, err = url.Parse(fmt.Sprintf("http://%s:%s@%s:%s", *pr.Username, *pr.Password, pr.IP, pr.Port))
		} else {
			proxyUrl, err = url.Parse(fmt.Sprintf("http://%s:%s", pr.IP, pr.Port))
		}

		if err != nil {
			log.Error("error parsing proxy", err)
			goto onErrorContinue
		}

		tk.Client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}

		tk.proxy = pr
		tk._proxyLocker.Unlock()

		if !firstCalculated {
			wg.Done()
			firstCalculated = true
		}

	onErrorContinue:
		tk._proxyLocker.Unlock()
		continue
	}
}

func (tk *BMonitor) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(tk.Ctx, method, url, body)
	if err != nil {
		log.Error("error creating request", err)
		return nil, err
	}


	if tk.proxy.Username != nil && tk.proxy.Password != nil && tk._proxyLocker.TryLockWithContext(tk.Ctx) {
		req.Header.Set("Proxy-Authorization", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", *tk.proxy.Username, *tk.proxy.Password))))
		tk._proxyLocker.Unlock()
	}
	return req, err
}
