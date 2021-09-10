package frame

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"os"
)

var rdb = ConnectRedis(os.Getenv("REDIS_URL"))

//SubscribeToChannel connects to a redis pub/sub stream and returns a pointer to a PubSub struct
func SubscribeToChannel(channelName string) (ps *PubSub, err error) {
	if len(channelName) == 0 {
		return nil, channelEmptyError
	}
	pubsub := rdb.Subscribe(context.Background(), channelName)

	_, err = pubsub.Receive(context.Background())
	if err != nil {
		return nil, err
	}
	return &PubSub{
		Channel: pubsub.Channel(),
		redisPS: pubsub,
	}, nil
}

//PubSub is a wrapper structure around redis.PubSub it provides the channel as a public field
type PubSub struct {
	Channel <-chan *redis.Message
	redisPS *redis.PubSub
}

//Close closes the underlying pub/sub stream as well as the channel attached to it
func (p *PubSub) Close() error {
	return p.redisPS.Close()
}

func (p *PubSub) Chan(ctx context.Context) <-chan map[string]interface{} {
	data := make(chan map[string]interface{})

	go func() {
		defer close(data)
		defer p.Close()
		for {
			select {
			case msg := <-p.redisPS.Channel():
				if msg == nil{
					continue
				}
				var dt map[string]interface{}
				if err := json.Unmarshal([]byte(msg.Payload), &dt); err != nil {
					continue
				}
				data <- dt
			case <-ctx.Done():
				return
			default:
				continue
			}
		}
	}()

	return data
}

func ConnectRedis(dsn string) *redis.Client {
	if !redisURLRegex.MatchString(dsn) {
		return nil
	}
	opts, err := redis.ParseURL(dsn)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opts)
}
