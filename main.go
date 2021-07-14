package main

import (
	"net/http"

	"vudung-practice-go/driver"
	"vudung-practice-go/src/modules/user/handler"
	"vudung-practice-go/src/modules/user/repoimpl"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	mongo := driver.ConnectMongoDB("mongodb+srv://vudung:vudung@cluster0.su9sk.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	userRepo := repoimpl.NewUserRepoMongo(mongo.Client.Database("go"))

	server := echo.New()
	// server.Use(middleware.Logger())
	// server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	//server.POST("/Login", controllers.Login
	IsLoggedIn := middleware.JWT([]byte("secret"))
	handler.NewUserHandler(server, userRepo, IsLoggedIn)
	server.Logger.Fatal(server.Start(":8080"))
}
