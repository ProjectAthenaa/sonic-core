package sonic

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var rdb = connectToRedis()

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

//connectToRedis is an internal method used to create a redis client to use with pub/sub
func connectToRedis() *redis.Client {
	redisURL := os.Getenv("REDIS_URL")

	if !redisURLRegex.MatchString(redisURL) {
		panic(redisFormatError)
	}

	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opts)
}
