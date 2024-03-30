package handlers

import (
	"cekzakat/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

var users = make(map[string]entity.User)

func RegisterHandler(c echo.Context) error {
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
	users[user.Username] = entity.User{
		Username: user.Username,
		Password: user.Password,
	}
	userResp.Message = "Berhasil Daftar"
	return c.JSON(http.StatusOK, userResp)
}

func LoginHandler(c echo.Context) error {
	user := entity.User{}
	userResp := entity.UserResponse{}

	if err := c.Bind(&user); err != nil {
		userResp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, userResp)
	}

	if _, found := users[user.Username]; !found {
		userResp.Message = "username tidak ada"
		return c.JSON(http.StatusInternalServerError, userResp)
	}

	if users[user.Username].Password != user.Password {
		userResp.Message = "password tidak ada"
		return c.JSON(http.StatusInternalServerError, userResp)
	}

	users[user.Username] = entity.User{
		Username: user.Username,
		Password: user.Password,
		IsLogin: true,
	}
	userResp.Message = "Berhasil Login"
	return c.JSON(http.StatusOK, userResp)

}

func LogoutHandler(c echo.Context) error {
	user := entity.User{}
	userResp := entity.UserResponse{}

	if err := c.Bind(&user); err != nil {
		userResp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, userResp)
	}

	if _, found := users[user.Username]; !found {
		userResp.Message = "username tidak ada"
		return c.JSON(http.StatusInternalServerError, userResp)
	}

	users[user.Username] = entity.User{
		Username: user.Username,
		Password: user.Password,
		IsLogin:  false,
	}
	userResp.Message = "berhasil logout"
	return c.JSON(http.StatusOK, userResp)

}

func IsLogin(username string) bool {
	if _, found := users[username]; !found {
		return false
	}
	return users[username].IsLogin
}
