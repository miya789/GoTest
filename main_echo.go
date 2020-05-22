package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// master
func main() {
	e := echo.New()

	// e.Use(middleware.Logger())
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":1234"))
	// fmt.Printf("start listening server at %s\n", "8080")
	// e.Start(fmt.Sprintf(":%v", 8080))
	// revert-test
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}
