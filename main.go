package main

import (
	"net/http"
	"os"

	"github.com/vudung18110263/Practice_Go/driver"
	"github.com/vudung18110263/Practice_Go/src/modules/user/handler"
	"github.com/vudung18110263/Practice_Go/src/modules/user/repoimpl"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	os := os.Getenv("PORT")
	mongo := driver.ConnectMongoDB("mongodb+srv://truongnv:1234@cluster0.2f1oc.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	userRepo := repoimpl.NewUserRepoMongo(mongo.Client.Database("go"))

	server := echo.New()
	// server.Use(middleware.Logger())
	// server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	//server.POST("/Login", controllers.Login
	server.GET("/", healthCheck)
	IsLoggedIn := middleware.JWT([]byte("secret"))
	handler.NewUserHandler(server, userRepo, IsLoggedIn)
	server.Logger.Fatal(server.Start(":" + os))
}
func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
