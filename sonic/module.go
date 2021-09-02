package sonic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Module struct {
	Name   string
	Fields []InputField
}

type InputField struct {
	Validation  string `json:"validation"`
	Label       string         `json:"label"`
	Property    string         `json:"property"`
	FormElement string         `json:"formElement"`
}

func RegisterModule(module *Module) error {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		return err
	}

	rdb := redis.NewClient(opts)

	val, err := json.Marshal(&module)
	if err != nil {
		return err
	}

	var ctx, cancel = context.WithCancel(context.Background())

	go func() {
		for range time.Tick(time.Second * 5) {
			rdb.SetNX(ctx, fmt.Sprintf("modules:%s", module.Name), string(val), redis.KeepTTL)
		}
	}()

	go func() {
		c := make(chan os.Signal, 1)
		defer close(c)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
		rdb.Del(ctx, fmt.Sprintf("modules:%s", module.Name))
	}()

	return nil
}
