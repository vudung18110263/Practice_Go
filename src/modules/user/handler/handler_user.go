package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	e.GET("/account/list", uh.FindAllHandler)
	e.PUT("/account", uh.UpdateUser)
	e.DELETE("/account", uh.Delete)
	e.GET("/account/:id/detail", uh.FindUser)
	e.POST("/account/SignIn", uh.SignIn)
	e.POST("/account/Login", uh.Login)
}

func (uh *UserHandler) RegisterHandler(c echo.Context) error {
	req := new(model.UserForm)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = uh.UserRepo.IsUserNameExist(req.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "username exist")
	}
	var user model.User
	user.Name = req.Name
	user.Password = req.Password
	err = uh.UserRepo.Insert(user)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, req)
}

func (uh *UserHandler) FindAllHandler(c echo.Context) error {

	var user []model.User

	//username := c.QueryParam("username")
	offset, err := strconv.Atoi(c.QueryParam("offset"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err = uh.UserRepo.FindAll(offset, limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	req := new(model.User)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var user model.User
	user.Id = req.Id
	user.Name = req.Name
	user.Password = req.Password

	err = uh.UserRepo.Update(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
func (uh *UserHandler) Delete(c echo.Context) error {
	req := new(model.DeleteUser)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = uh.UserRepo.Delete(req.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "delete successful")
}
func (uh *UserHandler) FindUser(c echo.Context) error {
	Id := c.Param("id")

	var err error
	var user *model.User

	user, err = uh.UserRepo.Find(Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
func (uh *UserHandler) SignIn(c echo.Context) error {
	req := new(model.UserForm)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var IsUser bool
	_, IsUser = uh.UserRepo.IsUser(req.Name, req.Password)
	if !IsUser {
		return c.JSON(http.StatusBadRequest, "false")
	}
	return c.JSON(http.StatusOK, true)
}

func (uh *UserHandler) Login(c echo.Context) error {
	req := new(model.UserForm)
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var IsUser bool
	_, IsUser = uh.UserRepo.IsUser(req.Name, req.Password)
	if !IsUser {
		return c.JSON(http.StatusBadRequest, "false")
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = req.Name
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &model.LoginResponse{
		Token: t,
	})
}
