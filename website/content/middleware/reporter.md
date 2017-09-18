+++
title = "Reporter Middleware"
description = "Reporter middleware for Echo. It provides error reporting using LabStack log service."
[menu.main]
  name = "Reporter"
  parent = "middleware"
+++


Reporter provides automatic error reporting and notification using the LabStack platform.

API key: https://labstack.com/signup<br>
Dashboard access: https://labstack.com/log

> Echo community contribution 

## Dependencies

```go
import (
  "github.com/labstack/echo-contrib/reporter"
)
```

*Usage*

```go
e := echo.New()
e.Use(reporter.Middleware("<ACCOUNT_ID>", "<API_KEY>"))
```

## Custom Configuration

*Usage*

```go
e := echo.New()
e.Use(reporter.MiddlewareWithConfig(reporter.Config{
  AccountID: "<ACCOUNT_ID>",
  APIKey: "<API_KEY>",
  Headers: []string{"User-Agent"}, // Headers to include
}))
```

## Configuration

```go
// Config defines the config for Reporter middleware.
Config struct {
  // Skipper defines a function to skip middleware.
  Skipper middleware.Skipper

  // App ID
  AppID string

  // App name
  AppName string

  // LabStack Account ID
  AccountID string `json:"account_id"`

  // LabStack API key
  APIKey string `json:"api_key"`

  // Headers to include
  Headers []string `json:"headers"`

  // TODO: To be implemented
  ClientLookup string `json:"client_lookup"`
}
```

[Learn more](https://labstack.com/docs/log)

*Default Configuration*

```go
// DefaultConfig is the default Reporter middleware config.
DefaultConfig = Config{
  Skipper: middleware.DefaultSkipper,
}
```

## Alert Policy

Navigate to https://labstack.com/monitor/new to create an alert policy with the following data:

- Name: <POLICY_NAME>
- Query String: app_name: <APP_NAME> AND level:FATAL
- Email: <EMAIL>
- Message:

```html
Message: {{ message }}
Stack Trace: {{ stack_trace }}
```

> Leave rest of the fields as is.
