package sonic

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	sonic "github.com/ProjectAthenaa/sonic-core/protos"
	"math/rand"
	"unsafe"
)

//ConvertProxyToString converts a struct of type sonic.Proxy to a string
func ConvertProxyToString(proxy *sonic.Proxy) string {
	var pr string

	if proxy.Username != nil && proxy.Password != nil {
		pr = fmt.Sprintf("http://%s:%s@%s:%s", *proxy.Username, *proxy.Password, proxy.IP, proxy.Port)
	} else {
		pr = fmt.Sprintf("http://%s:%s", proxy.IP, proxy.Password)
	}
	return pr
}

//GetRandomUserAgent retrieves a random user agent from the list of user agents
func GetRandomUserAgent() string {
	return userAgents[rand.Intn(len(userAgents)-1)]
}

//ErrString returns a pointer to the string error
func ErrString(err error) *string {
	e := err.Error()
	return &e
}

//UnsafeString returns a pointer to the byte slice with 0 allocations
func UnsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

type Map map[string]string

func (m Map) Value() (driver.Value, error) {
	return json.Marshal(&m)
}

func (m *Map) Scan(v interface{}) error {
	if v == nil {
		return nil
	}
	switch data := v.(type) {
	case string:
		return json.Unmarshal([]byte(data), &m)
	case []byte:
		return json.Unmarshal(data, &m)
	default:
		return fmt.Errorf("cannot scan type %t into Map", v)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}