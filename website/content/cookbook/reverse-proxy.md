+++
title = "Reverse Proxy"
description = "Using Echo as reverse proxy server"
draft = true
[menu.main]
  name = "Reverse Proxy"
  parent = "cookbook"
+++

## How to use Echo as a reverse proxy server?

### Step 1: Create a `struct` to store upstream servers

### Step 1: Create a HTTP proxy function using `httputil.NewSingleHostReverseProxy`

```go
func httpProxy(u *url.URL) http.Handler {
	return httputil.NewSingleHostReverseProxy(u)
}
```

### Step 1: [Generate a self-signed X.509 TLS certificate](/cookbook/http2#step-1-generate-a-self-signed-x-509-tls-certificate)

### Step 2: Register route to serve web assets

```go
e.Static("/", "static")
```

### Step 3: Create a handler to serve `index.html` and push it's dependencies

```go
e.GET("/", func(c echo.Context) (err error) {
  pusher, ok := c.Response().Writer.(http.Pusher)
  if ok {
    if err = pusher.Push("/app.css", nil); err != nil {
      return
    }
    if err = pusher.Push("/app.js", nil); err != nil {
      return
    }
    if err = pusher.Push("/echo.png", nil); err != nil {
      return
    }
  }
  return c.File("index.html")
})
```

If `http.Pusher` is supported, web assets are pushed; otherwise, client makes separate requests to get them.

### Step 4: Configure TLS server using `cert.pem` and `key.pem`

```go
e.StartTLS(":1323", "cert.pem", "key.pem")
```

### Step 5: Run the server and browse to https://localhost:1323

```
Protocol: HTTP/2.0
Host: localhost:1323
Remote Address: [::1]:60288
Method: GET
Path: /
```

## [Source Code]({{< source "http2-server-push" >}})

`index.html`

{{< embed "http2-server-push/index.html" >}}

`server.go`

{{< embed "http2-server-push/server.go" >}}

## Maintainers

- [vishr](https://github.com/vishr)
