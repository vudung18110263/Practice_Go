package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vudung18110263/Practice_Go/driver"
	"github.com/vudung18110263/Practice_Go/src/modules/user/handler"
	"github.com/vudung18110263/Practice_Go/src/modules/user/repository/repoimpl"
)

func main() {
	mongo := driver.ConnectMongoDB("mongodb://localhost:27017")
	userRepo := repoimpl.NewUserRepoMongo(mongo.Client.Database("go"))

	server := echo.New()

	//server.POST("/Login", controllers.Login)

	//IsLoggedIn := middleware.JWT([]byte("secret"))

	handler.NewUserHandler(server, userRepo)

	server.Logger.Fatal(server.Start(":8080"))
}
