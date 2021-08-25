# fasthttp TLS client

The fastest tls client out there

# How to use

- Create a new client

```go
    client := fasttls.NewClient(tls.HelloChrome_91, "http://username:password:ip:port")
  ```

- Add cookie jar to client

```go
    client.CreateCookieJar()
```

- Create and Make requests with headers and header order

```go
    req, err := client.NewRequest("POST", "https://example.com", []byte(`{"hello": "world"}`))

    headers := fasttls.Headers{
        "header1": {"value1"},
        "header2": {"value2"},
        "header3": {"value3"},
        fasttls.HeaderOrderKey: {
            "header3",
            "header2",
            "header1",
        },
    }

req.SetHeaders(headers) //or req.Headers = headers

resp, err := client.Do(req)
```

- Drop In Replacement with STD lib

```go
    client := fasttls.NewClientCompatibleWithStandardLibrary(tls.HelloChrome_91, "http://username:password:ip:port")
    req, err := client.NewRequest("POST", "https://example.com", strings.NewReader("some_test"))
    resp, err := client.Do(req)
    //resp is type *http.Response
```

- Destroy Client

```go
    client.Destroy()
```