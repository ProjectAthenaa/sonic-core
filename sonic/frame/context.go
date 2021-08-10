package frame

import (
	"errors"
	"github.com/ProjectAthenaa/sonic-core/sonic/database"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/go-redis/redis/v8"
	"sync"
)

type CoreContext struct {
	locker sync.Mutex
	store  sync.Map
}

var (
	errKeyExists  = errors.New("key already exists")
	errNotConnect = errors.New("rdb connect fail")
)

func (c *CoreContext) NewRedis(name string, dsn string) (redis.UniversalClient, error) {
	c.locker.Lock()
	defer c.locker.Unlock()

	if _, ok := c.store.Load(name); ok {
		return nil, errKeyExists
	}
	rds := ConnectRedis(dsn)
	if rds == nil {
		return nil, errNotConnect
	}
	c.store.Store(name, rds)
	return rds, nil
}

func (c *CoreContext) GetRedis(name string) redis.UniversalClient {
	if v, ok := c.store.Load(name); ok {
		return v.(redis.UniversalClient)
	}
	return nil
}

func (c *CoreContext) NewPg(name string, dsn string) (*ent.Client, error) {
	c.locker.Lock()
	defer c.locker.Unlock()

	if _, ok := c.store.Load(name); ok {
		return nil, errKeyExists
	}
	conn := database.Connect(dsn)
	if conn == nil {
		return nil, errNotConnect
	}
	c.store.Store(name, conn)
	return conn, nil
}

func (c *CoreContext) GetPg(name string) *ent.Client {
	if v, ok := c.store.Load(name); ok {
		return v.(*ent.Client)
	}
	return nil
}
