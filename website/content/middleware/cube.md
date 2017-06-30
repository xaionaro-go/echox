+++
date = "2014-09-09"
description = "Cube middleware for Echo. It provides HTTP analytics."
[menu.main]
  name = "Cube"
  parent = "middleware"
+++

Cube provides analytics for HTTP traffic. You can measure server latency, data transfer, discover top endpoints, top clients, slow requests and visualize key metrics such as total requests, client errors, server errors, status codes in a time series chart.

API key: https://labstack.com/signup<br>
Dashboard access: https://labstack.com/cube

## Dependencies

```go
import (
  "https://github.com/labstack/cube/echo"
)
```

*Usage*

```go
e := echo.New()
e.Use(cube.Middleware("<YOUR_API_KEY>"))
```

## Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(cube.MiddlewareWithConfig(cube.Config{
  APIKey: "<YOUR_API_KEY>",
  DispatchInterval: 5 * 60, // 5 minutes
}))
```

## Configuration

```go
// Config defines the config for Cube middleware.
Config struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Node name
  Node string `json:"node"`

  // Node group
  Group string `json:"group"`

  // LabStack API key
  APIKey string `json:"api_key"`

  // Number of requests in a batch
  BatchSize int `json:"batch_size"`

  // Interval in seconds to dispatch the batch
  DispatchInterval time.Duration `json:"dispatch_interval"`

  ClientLookup string `json:"client_lookup"`
}
```

[Learn more](https://labstack.com/docs/cube)

*Default Configuration*

```go
// DefaultConfig is the default Cube middleware config.
DefaultConfig = Config{
  Skipper: func(*http.Request) bool {
    return false
  },
  BatchSize:     60,
  DispatchInterval: 60,
}
```
