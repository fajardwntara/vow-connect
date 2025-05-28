package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fajardwntara/vow-connect/api/domain/user"
	"github.com/fajardwntara/vow-connect/helpers"
	"github.com/fajardwntara/vow-connect/utils"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userRepo user.UserRepository
}

func NewUserHandler(repo user.UserRepository) *userHandler {
	return &userHandler{userRepo: repo}
}

func (uh *userHandler) GetAllUsers(ctx *gin.Context) {

	users, err := uh.userRepo.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusNotFound, helpers.ErrorResponse{
			Success: false,
			Message: "data not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.SuccessResponse{
		Success: true,
		Data:    users,
	})
}

func (uh *userHandler) GetOne(ctx *gin.Context) {

	type UserResponse struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	userID := ctx.Param("id")

	parseUserID, err := strconv.ParseUint(userID, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: "validation error",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	user, err := uh.userRepo.GetByID(ctx.Request.Context(), uint(parseUserID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: "validation error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	ctx.JSON(http.StatusOK, helpers.SuccessResponse{
		Success: true,
		Message: "user found",
		Data:    user,
	})
}

func (uh *userHandler) Create(ctx *gin.Context) {

	var payload user.User
	type UserResponse struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: "validation error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	user := &user.User{
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = uh.userRepo.Create(ctx.Request.Context(), user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: "validation error",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	ctx.JSON(http.StatusCreated, helpers.SuccessResponse{
		Success: true,
		Message: "the user has successfully has been created",
		Data: &UserResponse{
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: utils.FormatDateTimeINA(user.CreatedAt),
			UpdatedAt: utils.FormatDateTimeINA(user.CreatedAt),
		},
	})
}

func (uh *userHandler) Delete(ctx *gin.Context) {

	userID := ctx.Param("id")

	// parse user id to uint
	parseUserID, err := strconv.ParseUint(userID, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: "validation error",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	// check whether the user id exist
	user, err := uh.userRepo.GetByID(ctx.Request.Context(), uint(parseUserID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: "data not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	err = uh.userRepo.Delete(ctx.Request.Context(), uint(user.ID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: "validation error",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	ctx.JSON(http.StatusOK, helpers.SuccessResponse{
		Success: true,
		Message: fmt.Sprintf("The user with username: %s has succesfully been deleted.", user.Username),
	})
}
