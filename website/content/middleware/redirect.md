+++
title = "Redirect"
[menu.side]
  name = "Redirect"
  parent = "middleware"
  weight = 5
+++

## HTTPSRedirect Middleware

HTTPSRedirect middleware redirects HTTP requests to HTTPS.
For example, http://labstack.com will be redirect to https://labstack.com.

*Usage*

```go
e := echo.New()
e.Pre(middleware.HTTPSRedirect())
```

## HTTPSWWWRedirect Middleware

HTTPSWWWRedirect redirects HTTP requests to WWW HTTPS.
For example, http://labstack.com will be redirect to https://www.labstack.com.

*Usage*

```go
e := echo.New()
e.Pre(middleware.HTTPSWWWRedirect())
```

## WWWRedirect Middleware

WWWRedirect redirects non WWW requests to WWW.

For example, http://labstack.com will be redirect to http://www.labstack.com.

*Usage*

```go
e := echo.New()
e.Pre(middleware.WWWRedirect())
```

## NonWWWRedirect Middleware

NonWWWRedirect redirects WWW request to non WWW.
For example, http://www.labstack.com will be redirect to http://labstack.com.

*Usage*

```go
e := echo.New()
e.Pre(middleware.NonWWWRedirect())
```

### Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(middleware.HTTPSRedirectWithConfig(middleware.RedirectConfig{
  Code: http.StatusTemporaryRedirect,
}))
```

This will redirect the request HTTP to HTTPS with status code `307 - StatusTemporaryRedirect`.

### Configuration

```go
RedirectConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Status code to be used when redirecting the request.
  // Optional. Default value http.StatusMovedPermanently.
  Code int `json:"code"`
}
```

*Default Configuration*

```go
DefaultRedirectConfig = RedirectConfig{
  Skipper: defaultSkipper,
  Code:    http.StatusMovedPermanently,
}
```
