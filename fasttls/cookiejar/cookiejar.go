package cookiejar

import (
	"io"
	"sync"
	"time"
	"unsafe"

	"github.com/ProjectAthenaa/sonic-core/fasttls/fasthttp"
)

var jarMutexes = sync.Map{}

var cookiePool = sync.Pool{
	New: func() interface{} {
		return &CookieJar{}
	},
}

// AcquireCookieJar returns an empty CookieJar object from pool
func AcquireCookieJar() *CookieJar {
	cj := cookiePool.Get().(*CookieJar)
	jarMutexes.Store(cj, &sync.Mutex{})
	return cj
}

// ReleaseCookieJar returns CookieJar to the pool
func ReleaseCookieJar(c *CookieJar) {
	mu, _ := jarMutexes.LoadAndDelete(c)
	mu.(*sync.Mutex).Lock()
	defer mu.(*sync.Mutex).Unlock()
	c.Release()
	cookiePool.Put(c)
}

// CookieJar is container of cookies
//
// This object is used to handle multiple cookies
type CookieJar map[string]*fasthttp.Cookie

// Set sets cookie using key-value
//
// This function can replace an existent cookie
func (cj *CookieJar) Set(key, value string) {
	setCookie(cj, key, value)
}

// Get returns and delete a value from cookiejar.
func (cj *CookieJar) Get() *fasthttp.Cookie {
	for k, v := range *cj {
		delete(*cj, k)
		return v
	}
	return nil
}

// SetBytesK sets cookie using key=value
//
// This function can replace an existent cookie.
func (cj *CookieJar) SetBytesK(key []byte, value string) {
	setCookie(cj, b2s(key), value)
}

// SetBytesV sets cookie using key=value
//
// This function can replace an existent cookie.
func (cj *CookieJar) SetBytesV(key string, value []byte) {
	setCookie(cj, key, b2s(value))
}

// SetBytesKV sets cookie using key=value
//
// This function can replace an existent cookie.
func (cj *CookieJar) SetBytesKV(key, value []byte) {
	setCookie(cj, b2s(key), b2s(value))
}

func setCookie(cj *CookieJar, key, value string) {
	cj.lock()
	defer cj.unlock()
	c, ok := (*cj)[key]
	if !ok {
		c = fasthttp.AcquireCookie()
	}
	c.SetKey(key)
	c.SetValue(value)
	(*cj)[key] = c
}

// SetCookie sets cookie using its key.
//
// After that you can use Peek or Get function to get cookie value.
func (cj *CookieJar) Put(cookie *fasthttp.Cookie) {
	cj.lock()
	defer cj.unlock()
	c, ok := (*cj)[b2s(cookie.Key())]
	if ok {
		fasthttp.ReleaseCookie(c)
	}
	(*cj)[b2s(cookie.Key())] = cookie
}

// Peek peeks cookie value using key.
//
// This function does not delete cookie
func (cj *CookieJar) Peek(key string) *fasthttp.Cookie {
	return (*cj)[key]
}

// Release releases all cookie values.
func (cj *CookieJar) Release() {
	for k := range *cj {
		cj.ReleaseCookie(k)
	}
}

// ReleaseCookie releases a cookie specified by parsed key.
func (cj *CookieJar) ReleaseCookie(key string) {
	c, ok := (*cj)[key]
	if ok {
		fasthttp.ReleaseCookie(c)
		delete(*cj, key)
	}
}

// PeekValue returns value of specified cookie-key.
func (cj *CookieJar) PeekValue(key string) []byte {
	c, ok := (*cj)[key]
	if ok {
		return c.Value()
	}
	return nil
}

// ReadResponse gets all Response cookies reading Set-Cookie header.
func (cj *CookieJar) ReadResponse(r *fasthttp.Response, domain string) {
	r.Header.VisitAllCookie(func(key, value []byte) {
		cookie := fasthttp.AcquireCookie()
		cookie.ParseBytes(value)
		cookie.SetDomain(domain)
		cj.Put(cookie)
	})
}

// ReadRequest gets all cookies from a Request reading Set-Cookie header.
func (cj *CookieJar) ReadRequest(r *fasthttp.Request) {
	r.Header.VisitAllCookie(func(key, value []byte) {
		cookie := fasthttp.AcquireCookie()
		cookie.ParseBytes(value)
		cj.Put(cookie)
	})
}

// WriteTo writes all cookies representation to w.
func (cj *CookieJar) WriteTo(w io.Writer) (n int64, err error) {
	for _, c := range *cj {
		nn, err := c.WriteTo(w)
		n += nn
		if err != nil {
			break
		}
	}
	return
}

// FillRequest dumps all cookies stored in cj into Request adding this values to Cookie header.
func (cj *CookieJar) FillRequest(r *fasthttp.Request) {
	for _, c := range *cj {
		expires := c.Expire()
		if !expires.IsZero() {
			if expires.Unix() < time.Now().Unix() {
				continue
			}
		}

		//cookieDomain := string(c.Domain())
		//fullHost := string(r.Host())
		//u, err := tld.Parse(string(r.URI().FullURI()))
		//parentHost := ""
		//if err == nil {
		//	parentHost = fmt.Sprintf("%s.%s", u.Domain, u.TLD)
		//	fullHost = parentHost
		//	if u.Subdomain != "" {
		//		fullHost = fmt.Sprintf("%s.%s.%s", u.Subdomain, u.Domain, u.TLD)
		//	}
		//}
		//
		//allowedCookie := false
		//if fullHost == cookieDomain {
		//	allowedCookie = true
		//} else if parentHost != "" {
		//	if strings.HasSuffix(parentHost, cookieDomain) || strings.HasSuffix("."+parentHost, cookieDomain) {
		//		allowedCookie = true
		//	}
		//}
		//if allowedCookie {
		r.Header.SetCookieBytesKV(c.Key(), c.Value())
		//}
	}
}

// FillResponse dumps all cookies stored in cj into Response adding this values to Cookie header.
func (cj *CookieJar) FillResponse(r *fasthttp.Response) {
	for _, c := range *cj {
		r.Header.SetCookie(c)
	}
}

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func (cj *CookieJar) lock() {
	mu, ok := jarMutexes.Load(cj)
	if !ok {
		mu = &sync.Mutex{}
		jarMutexes.Store(cj, mu)
	}
	mu.(*sync.Mutex).Lock()
}

func (cj *CookieJar) unlock() {
	mu, ok := jarMutexes.Load(cj)
	if !ok {
		mu = &sync.Mutex{}
		jarMutexes.Store(cj, mu)
		return
	}
	mu.(*sync.Mutex).Unlock()
}
