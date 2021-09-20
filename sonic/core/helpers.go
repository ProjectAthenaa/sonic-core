package core

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/logs"
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/accountgroup"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/product"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/task"
	"os"
	"sort"
	"strings"
	"sync"
)

type scratchTask struct {
	ctx context.Context
	*ent.Task
}

func getPayload(ctx context.Context, taskID string) (*module.Data, error) {
	dbTask, err := Base.GetPg("pg").
		Task.
		Query().
		WithProfileGroup().
		WithProduct().
		WithProxyList().
		WithTaskGroup().
		Where(
			task.ID(
				sonic.UUIDParser(taskID),
			),
		).
		First(ctx)
	if err != nil {
		log.Errorf("[server] [error retrieving task] [%s] [%s]", fmt.Sprint(err), taskID)
		return nil, err
	}

	tsk := scratchTask{
		ctx:  ctx,
		Task: dbTask,
	}

	//go tsk.processMonitor(ctx)

	var mData *module.Data

	prod := tsk.Edges.Product[0]

	mData = &module.Data{
		TaskData: &module.TaskData{
			Color: prod.Colors,
			Size:  prod.Sizes,
		},
		Channels: &module.Channels{
			MonitorChannel:  tsk.getMonitorID(),
			UpdatesChannel:  tsk.ID.String(),
			CommandsChannel: hash(tsk.ID.String()),
		},
	}

	if len(prod.Colors) == 0 || prod.Colors[0] == "random" {
		mData.TaskData.Color = []string{"0"}
		mData.TaskData.RandomColor = true
	}

	if len(prod.Sizes) == 0 || prod.Sizes[0] == "random" {
		mData.TaskData.Size = []string{"0"}
		mData.TaskData.RandomSize = true
	}

	var wg = &sync.WaitGroup{}

	go func() {
		wg.Add(1)
		defer wg.Done()
		mData.Profile, _ = tsk.getProfile()
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		mData.Proxy, _ = tsk.getProxy()
	}()

	mData.Metadata = prod.Metadata
	if siteNeedsAccount[tsk.Edges.Product[0].Site] {
		go func() {
			wg.Add(1)
			defer wg.Done()
			mData.Metadata["username"], mData.Metadata["password"], _ = tsk.getAccount()
		}()
	}

	mData.TaskID = tsk.ID.String()
	wg.Wait()

	return mData, nil
}

func (j *scratchTask) getProxy() (*module.Proxy, error) {
	dbProxyList := j.Edges.ProxyList[0]

	dbProxies, err := dbProxyList.Proxies(j.ctx)
	if err != nil {
		log.Errorf("[server] [error loading proxies] [%s] [%s]", j.ID.String(), fmt.Sprint(err))
		return nil, err
	}

	key := fmt.Sprintf("tasks:proxies:%s", dbProxyList.ID.String())

	locker := redisSync.NewMutex(key + ":locker")

	if err = locker.LockContext(j.ctx); err != nil {
		log.Errorf("[server] [error acquiring proxy mutex] [%s] [%s]", j.ID.String(), fmt.Sprint(err))
		return nil, err
	}

	defer func() {
		if ok, err := locker.UnlockContext(j.ctx); !ok || err != nil {
			log.Error("error unlocking proxy mutex: ", err)
		}
	}()

	proxies := rdb.SMembers(j.ctx, key).Val()

	if len(proxies) == 0 {
		var availablePool []interface{}

		for _, proxy := range dbProxies {
			var payload []byte
			if payload, err = json.Marshal(&proxy); err != nil {
				continue
			}
			availablePool = append(availablePool, string(payload))
		}

		rdb.SAdd(j.ctx, key, availablePool[1:]...)

		return &module.Proxy{
			Username: &dbProxies[0].Username,
			Password: &dbProxies[0].Password,
			IP:       dbProxies[0].IP,
			Port:     dbProxies[0].Port,
		}, nil
	}

	var proxy *module.Proxy

	data := rdb.SPop(j.ctx, key).Val()

	if err = json.Unmarshal([]byte(data), &proxy); err != nil {
		log.Errorf("[server] [error unmarshalling proxy] [%s] [%s]", j.ID.String(), fmt.Sprint(err))
		return nil, err
	}

	return proxy, nil
}

func (j *scratchTask) getAccount() (username, password string, err error) {

	app, _ := j.Edges.TaskGroup.App(j.ctx)
	dbAccounts, err := app[0].
		QueryAccountGroups().
		Where(
			accountgroup.SiteEQ(
				accountgroup.Site(
					j.Edges.Product[0].Site,
				),
			),
		).
		First(j.ctx)
	if err != nil {
		log.Errorf("[server] [error retrieving accounts] [%s] [%s]", j.ID.String(), fmt.Sprint(err))
		return "", "", err
	}

	key := fmt.Sprintf("tasks:accounts:%s", dbAccounts.ID.String())

	locker := redisSync.NewMutex(key + ":locker")

	if err = locker.LockContext(j.ctx); err != nil {
		log.Errorf("[server] [error acquiring account group mutex] [%s] [%s]", j.ID.String(), fmt.Sprint(err))
		return "", "", err
	}

	defer func() {
		if ok, err := locker.UnlockContext(j.ctx); !ok || err != nil {
			log.Errorf("[server] [error unlocking account group mutex] [%s] [%s]", j.ID.String(), fmt.Sprint(err))
			log.Error("error unlocking account group mutex: ", err)
		}
	}()

	accounts := rdb.SMembers(j.ctx, key).Val()

	if len(accounts) == 0 {
		var availablePool []interface{}

		for u, p := range dbAccounts.Accounts {
			availablePool = append(availablePool, fmt.Sprintf("%s:%s", u, p))
		}

		rdb.SAdd(j.ctx, key, availablePool[1:]...)

		acc := strings.Split(availablePool[0].(string), ":")

		return acc[0], acc[1], nil
	}

	data := rdb.SPop(j.ctx, key).Val()

	acc := strings.Split(data, ":")

	return acc[0], acc[1], nil
}

func (j *scratchTask) getProfile() (retProf *module.Profile, err error) {
	profileGroup := j.Edges.ProfileGroup

	key := fmt.Sprintf("tasks:profiles:%s", profileGroup.ID.String())

	locker := redisSync.NewMutex(key + ":locker")

	if err := locker.LockContext(j.ctx); err != nil {
		log.Errorf("[server] [error acquiring profile group mutex] [%s] [%s]", j.ID.String(), fmt.Sprint(err))
		return nil, err
	}

	defer func() {
		if ok, err := locker.UnlockContext(j.ctx); !ok || err != nil {
			log.Errorf("[server] [error unlocking profile group mutex] [%s] [%s]", j.ID.String(), fmt.Sprint(err))
		}
	}()

	accounts := rdb.SMembers(j.ctx, key).Val()

	if len(accounts) == 0 {
		var availablePool []interface{}
		var toAppend *module.Profile

		profiles, err := profileGroup.Profiles(j.ctx)
		if err != nil {
			return nil, sonic.EntErr(err)
		}

		for i, prof := range profiles {
			shipping, err := prof.QueryShipping().First(j.ctx)
			if err != nil {
				return nil, sonic.EntErr(err)
			}

			shippingAddress, err := shipping.QueryShippingAddress().First(j.ctx)
			if err != nil {
				return nil, sonic.EntErr(err)
			}

			billing, err := prof.QueryBilling().First(j.ctx)
			if err != nil {
				return nil, sonic.EntErr(err)
			}

			if os.Getenv("DEBUG") == "1" {
				log.Info("Debug enabled, accessing profiles without encryption")
			} else {
				billing = billing.Decrypt()
			}

			toAppend = &module.Profile{
				Email: prof.Email,
				Shipping: &module.Shipping{
					FirstName:   shipping.FirstName,
					LastName:    shipping.LastName,
					PhoneNumber: shipping.PhoneNumber,
					ShippingAddress: &module.Address{
						AddressLine:  shippingAddress.AddressLine,
						AddressLine2: &shippingAddress.AddressLine2,
						Country:      shippingAddress.Country,
						State:        shippingAddress.State,
						City:         shippingAddress.City,
						ZIP:          shippingAddress.ZIP,
					},
					BillingAddress:    nil,
					BillingIsShipping: shipping.BillingIsShipping,
				},
				Billing: &module.Billing{
					Number:          billing.CardNumber,
					ExpirationMonth: billing.ExpiryMonth,
					ExpirationYear:  billing.ExpiryYear,
					CVV:             billing.CVV,
				},
			}

			if !shipping.BillingIsShipping {
				billingAddress, err := shipping.QueryBillingAddress().First(j.ctx)
				if err != nil {
					log.Error("query billing address: ", err)
					panic(err)
				}
				toAppend.Shipping.BillingAddress = &module.Address{
					AddressLine:  billingAddress.AddressLine,
					AddressLine2: &billingAddress.AddressLine2,
					Country:      billingAddress.Country,
					State:        billingAddress.State,
					City:         billingAddress.City,
					ZIP:          billingAddress.ZIP,
				}
			}

			payload, err := json.Marshal(&toAppend)
			if err != nil {
				return nil, err
			}

			if i == 0 {
				retProf = toAppend
			}

			availablePool = append(availablePool, string(payload))
		}

		rdb.SAdd(j.ctx, key, availablePool[1:]...)

		return retProf, nil
	}

	data := rdb.SPop(j.ctx, key).Val()

	var prof *module.Profile

	if err := json.Unmarshal([]byte(data), &prof); err != nil {
		log.Errorf("[server] [error unmarshalling profile] [%s] [%s]", j.ID.String(), fmt.Sprint(err))
		return nil, err
	}

	return prof, nil
}

func (j *scratchTask) getMonitorID() string {
	prefix := fmt.Sprintf("monitors:%s:", j.Edges.Product[0].Site)
	v := j.Edges.Product[0]
	switch v.LookupType {
	case product.LookupTypeLink:
		return prefix + hash(v.Link)
	case product.LookupTypeKeywords:
		sort.Strings(v.PositiveKeywords)
		sort.Strings(v.NegativeKeywords)

		for i, s := range v.PositiveKeywords {
			v.PositiveKeywords[i] = strings.ToLower(s)
		}
		for i, s := range v.NegativeKeywords {
			v.NegativeKeywords[i] = strings.ToLower(s)
		}

		return prefix + hash(strings.Join(v.PositiveKeywords, "")+strings.Join(v.NegativeKeywords, ""))
	case product.LookupTypeOther:
		for k, val := range v.Metadata {
			if strings.Contains(k, "LOOKUP_") {
				return prefix + hash(val)
			}
		}
	}
	return ""
}

func hash(text string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
