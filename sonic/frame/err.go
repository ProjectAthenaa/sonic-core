package frame

import (
	"errors"
	"regexp"
)

var channelEmptyError = errors.New("channel_name_cannot_be_empty")
var (
	redisURLRegex    = regexp.MustCompile(`rediss://\w+:\w+@.*:\d+`)
	redisFormatError = errors.New("redis address needs to have correct format")
)
