+++
title = "Proxy Middleware"
description = "Reverse proxy middleware for Echo"
[menu.main]
  name = "Proxy"
  parent = "middleware"
  weight = 6
+++

Proxy provides an HTTP/WebSocket reverse proxy middleware. It forwards a request
to upstream server using a configured load balancing technique.

*Usage*

```go
e.Use(middleware.Proxy(middleware.ProxyConfig{
  Targets: []*ProxyTarget{
    &ProxyTarget{
      URL: "http://t1",
    },
    &ProxyTarget{
      URL: "http://t2",
    },
  },
}))
```

## Configuration

```go
// ProxyConfig defines the config for Proxy middleware.
ProxyConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Load balancing technique.
  // Optional. Default value "random".
  // Possible values:
  // - "random"
  // - "round-robin"
  Balance string `json:"balance"`

  // Upstream target URLs
  // Required.
  Targets []*ProxyTarget `json:"targets"`
```

*Default Configuration*

Name | Value
---- | -----
Skipper | DefaultSkipper
Balance | "random"
