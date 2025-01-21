package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IuserLikeUsecase interface {
	GetUsers(userId uint) (model.UserResponse, error)
}

type userLikeUsecase struct {
	ulr repository.IUserLikeRepository
	// uv validator.IUserValidator
}

func NewUserLikeUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IuserLikeUsecase {
	return &userLikeUsecase{ur, uv}
}

func (ulu *userLikeUsecase) GetUsers(userId uint) (model.UserResponse, error) {
	Users := model.User{}
	UserLike := model.UserLike{}
	if err := uu.ur.GetSelf(&Users, userId); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    loginUser.ID,
		Email: loginUser.Email,
		Name:  loginUser.Name,
		Age:   loginUser.Age,
	}
	return resUser, nil
}
