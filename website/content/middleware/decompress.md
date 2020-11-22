+++
title = "Decompress Middleware"
description = "Decompress middleware for Echo"
[menu.main]
  name = "Decompress"
  parent = "middleware"
+++

Decompress middleware decompresses HTTP request if Content-Encoding header is set to gzip.

*Usage*

`e.Use(middleware.Decompress())`

## Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.DecompressWithConfig(middleware.DecompressConfig{
  Skipper: Skipper
}))
```

## Configuration

```go
DecompressConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper
}
```

*Default Configuration*

```go
DefaultDecompressConfig = DecompressConfig{
  Skipper: DefaultSkipper,
}
```
