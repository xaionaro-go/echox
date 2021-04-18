package main

import (
	"crypto/rand"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aead/chacha20"
	"github.com/aead/chacha20/chacha"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	myEncryptionKey = make([]byte, chacha.KeySize)
	mySigningKey    = []byte("mySigningKey")
)

func init() {
	copy(myEncryptionKey, `myEncryptionKey`)
}

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtCustomClaimsEncrypted jwtCustomClaims

type encryptionContainer struct {
	IV        []byte
	Encrypted []byte
}

func (claims *jwtCustomClaimsEncrypted) MarshalJSON() ([]byte, error) {
	decrypted, err := json.Marshal((*jwtCustomClaims)(claims))
	if err != nil {
		return nil, err
	}
	encrypted := make([]byte, len(decrypted))
	iv := make([]byte, 24)
	rand.Read(iv)
	chacha20.XORKeyStream(encrypted, decrypted, iv, myEncryptionKey)
	return json.Marshal(encryptionContainer{IV: iv, Encrypted: encrypted})
}

func (claims *jwtCustomClaimsEncrypted) UnmarshalJSON(encrypted []byte) error {
	container := encryptionContainer{}
	err := json.Unmarshal(encrypted, &container)
	if err != nil {
		return err
	}
	decrypted := make([]byte, len(container.Encrypted))
	chacha20.XORKeyStream(decrypted, container.Encrypted, container.IV, myEncryptionKey)
	return json.Unmarshal(decrypted, (*jwtCustomClaims)(claims))
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	expiresAt := time.Now().Add(time.Hour * 72)
	claims := &jwtCustomClaimsEncrypted{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(mySigningKey)
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:    "X-Jwt",
		Value:   t,
		Path:    "/",
		Expires: expiresAt,
	})

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func loginForm(c echo.Context) error {
	return c.HTML(200, `
<html>
	<head></head>
	<body>
		<form method="POST" action="/login">
			<input name="username" type="text"     placeholder="username"><br>
			<input name="password" type="password" placeholder="password"><br>
			<input type="submit">
		</form>
	</body>
</html>`)
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaimsEncrypted)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", loginForm)

	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:      &jwtCustomClaimsEncrypted{},
		SigningKey:  mySigningKey,
		TokenLookup: "cookie:X-Jwt",
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":1323"))
}
