+++
title = "Rewrite Middleware"
description = "Rewrite middleware for Echo"
[menu.main]
  name = "Rewrite"
  parent = "middleware"
+++

Rewrite middleware rewrites the URL path based on provided rules. It can be helpful for backward compatibility or just creating cleaner and more descriptive links.

*Usage*

```go
e.Pre(middleware.Rewrite(map[string]string{
  "/old":              "/new",
  "/api/*":            "/$1",
  "/js/*":             "/public/javascripts/$1",
  "/users/*/orders/*": "/user/$1/order/$2",
}))
```

The values captured in asterisk can be retrieved by index e.g. $1, $2 and so on.

## Custom Configuration

*Usage*

```go
e := echo.New()
e.Pre(middleware.RewriteWithConfig(middleware.RewriteConfig{}))
```

## Configuration

```go
// RewriteConfig defines the config for Rewrite middleware.
RewriteConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Rules defines the URL path rewrite rules.
  Rules map[string]string `yaml:"rules"`
}
```

*Default Configuration*

Name | Value
---- | -----
Skipper | DefaultSkipper

> Rewrite middleware should be registered via `Echo#Pre()` to get triggered before the router.
