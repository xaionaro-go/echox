package main

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	users = []string{"Joe", "Veer", "Zion"}
)

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

// allowOrigin takes the origin as an argument and returns true if the origin
// is allowed or false otherwise.
func allowOrigin(origin string) (bool, error) {
	// In this example we use a regular expression but we can imagine various
	// kind of custom logic. For example, an external datasource could be used
	// to maintain the list of allowed origins.
	return regexp.MatchString(`^https:\/\/labstack\.(net|com)$`, origin)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS restricted with a custom function to allow origins
	// and with the GET, PUT, POST or DELETE methods allowed.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: allowOrigin,
		AllowMethods:    []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/api/users", getUsers)

	e.Logger.Fatal(e.Start(":1323"))
}
