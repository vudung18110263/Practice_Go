package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vudung18110263/Practice_Go/src/modules/user/model"
	"github.com/vudung18110263/Practice_Go/src/modules/user/repository"
)

type UserHandler struct {
	UserRepo repository.UserRepository
}

func NewUserHandler(e *echo.Echo, ur repository.UserRepository) {
	uh := &UserHandler{
		UserRepo: ur,
	}
	e.POST("/account", uh.RegisterHandler)
	// e.GET("/users/:userID", uh.GetUser)
	// e.GET("/users", uh.GetAllUser)

}

func (uh *UserHandler) RegisterHandler(c echo.Context) error {
	req := new(model.UserForm)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var user model.User
	user.Name = req.Name
	user.Password = req.Password
	err = uh.UserRepo.Insert(user)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusConflict, err.Error())
	}
	return c.JSON(http.StatusCreated, req)
}
