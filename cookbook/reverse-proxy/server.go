package main

type (
	Proxy struct {
		Targets []string
	}
)

func main() {
	e := echo.New()
	e.Use
	p := &Proxy{
		Targets: []string{
			"http://localhost:8081",
			"http://localhost:8082",
		},
	}
}
