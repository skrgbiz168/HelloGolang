package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IUserLikeController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
	GetSelf(c echo.Context) error
}

type userLikeController struct {
	uu usecase.IUserUsecase
}

func NewUserLikeController(uu usecase.IUserUsecase) IUserController {
	return &userLikeController{uu}
}

func (uc *userController) GetUsers (c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	usersRes, err := ulc.ulu.GetUsers(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, usersRes)
}

func (uc *userController) GetSelf(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	userRes, err := uc.uu.GetSelf(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userRes)
	// return c.JSON(http.StatusOK, echo.Map{
	// 	"email": userRes.Email,
	// 	"name":  userRes.Name,
	// 	"age":   userRes.Age,
	// })
	

}
