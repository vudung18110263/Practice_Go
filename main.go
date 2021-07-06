package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vudung18110263/Practice_Go/controllers"
)

func main() {
	server := echo.New()

	server.POST("/Login", controllers.Login)
	server.GET("/", hello)

	server.Logger.Fatal(server.Start(":8080"))
}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello word")
}
