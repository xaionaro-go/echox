+++
title = "Proxy Middleware"
description = "Reverse proxy middleware for Echo"
[menu.main]
  name = "Proxy"
  parent = "middleware"
+++

Proxy provides an HTTP/WebSocket reverse proxy middleware. It forwards a request
to upstream server using a configured load balancing technique.

*Usage*

```go
url1, err := url.Parse("http://localhost:8081")
if err != nil {
  e.Logger.Fatal(err)
}
url2, err := url.Parse("http://localhost:8082")
if err != nil {
  e.Logger.Fatal(err)
}
e.Use(middleware.Proxy(&middleware.RoundRobinBalancer{
  Targets: []*middleware.ProxyTarget{
    &middleware.ProxyTarget{
      URL: url1,
    },
    &middleware.ProxyTarget{
      URL: url2,
    },
  },
}))
```

## Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{}))
```

## Configuration

```go
// ProxyConfig defines the config for Proxy middleware.
ProxyConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Balancer defines a load balancing technique.
  // Required.
  // Possible values:
  // - RandomBalancer
  // - RoundRobinBalancer
  Balancer ProxyBalancer
}
```

*Default Configuration*

Name | Value
---- | -----
Skipper | DefaultSkipper

## [Example]({{< ref "cookbook/reverse-proxy.md">}})
