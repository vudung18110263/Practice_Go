package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vudung18110263/Practice_Go/entities"
)

func Login(c echo.Context) error {
	req := new(entities.Login)
	c.Bind(req)
	if req.Name != "admin" || req.Password != "123" {
		return c.JSON(http.StatusUnauthorized, "Invalid")
	}
	return c.JSON(http.StatusOK, "ok")
}
