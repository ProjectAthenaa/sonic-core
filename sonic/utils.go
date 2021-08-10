package sonic

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	sonic "github.com/ProjectAthenaa/sonic-core/protos"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"io"
	"io/ioutil"
	"math/rand"
	"net/url"
	"strings"
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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func UUIDParser(id string) uuid.UUID {
	parsedID, _ := uuid.Parse(id)
	return parsedID
}

func ErrorContains(err error, substring string) bool {
	return strings.Contains(fmt.Sprint(err), substring)
}

func IPFromContext(ctx context.Context) (ip string) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		return md.Get("x-real-ip")[0]
	}

	p, _ := peer.FromContext(ctx)
	if p != nil {
		return p.Addr.String()
	}

	return "Unknown"
}

func NopCloserBody(b io.ReadCloser) (io.ReadCloser, []byte) {
	body, _ := ioutil.ReadAll(b)
	return ioutil.NopCloser(bytes.NewBuffer(body)), body
}

func GetMonitorProxy(site Site) (proxy *url.URL, authorization string, err error) {
	url, err := url.Parse("http://KJND3:5Z6GNXPD@45.84.101.178:7249")
	if err != nil {
		return nil, "", err
	}

	return url, base64.StdEncoding.EncodeToString([]byte("KJND3:5Z6GNXPD")), nil
}

func GrabValueFromHTMLName(name string, html *[]byte) string {
	return strings.ReplaceAll(strings.Split(strings.Split(strings.Split(string(*html), name)[1], "/>")[0], "value=")[1], "\"", "")
}
