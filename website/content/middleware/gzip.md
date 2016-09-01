+++
title = "Gzip"
[menu.side]
  name = "Gzip"
  parent = "middleware"
  weight = 5
+++

## Gzip Middleware

Gzip middleware compresses HTTP response using gzip compression scheme.

*Usage*

`e.Use(middleware.Gzip())`

### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
  Level: 5,
}))
```

### Configuration

```go
GzipConfig struct {
  // Gzip compression level.
  // Optional. Default value -1.
  Level int
}
```

*Default Configuration*

```go
DefaultGzipConfig = GzipConfig{
  Level: -1,
}
```
