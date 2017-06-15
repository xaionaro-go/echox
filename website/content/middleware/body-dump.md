+++
title = "Body Dump Middleware"
description = "Body dump middleware for Echo"
[menu.main]
  name = "Body Dump"
  parent = "middleware"
  weight = 4
+++

Body dump middleware captures the request and response payload and calls the registered handler.

*Usage*

```go
e := echo.New()
e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte {
})))
```

## Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{}))
```

## Configuration

```go
BodyDumpConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Handler receives request and response payload.
  // Required.
  Handler BodyDumpHandler
}
```

*Default Configuration*

```go
DefaultBodyDumpConfig = BodyDumpConfig{
  Skipper: DefaultSkipper,
}
```
