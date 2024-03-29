package handlers

import (
	"cekzakat/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)


var users =  make(map[string]string)

func RegisterHandler(c echo.Context) error{

	user := entity.User{}
	userResp := entity.UserResponse{}
	if err := c.Bind(&user); err != nil {
		userResp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, userResp)
	}

	if _, found := users[user.Username]; found {
		userResp.Message = "username sudah terdaftar"
		return c.JSON(http.StatusInternalServerError, userResp)
	}

	users[user.Username] = user.Password
	userResp.Message = "Berhasil Daftar"
	return c.JSON(http.StatusOK, userResp)
}