package auth

import (
	"net/http"
	"time"

	"github.com/fajardwntara/vow-connect/api/domain/user"
	"github.com/fajardwntara/vow-connect/helpers"
	"github.com/fajardwntara/vow-connect/pkg/database"
	"github.com/fajardwntara/vow-connect/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {

	var userRequest struct {
		Username string
		Password string
	}

	type userPayload struct {
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Token     string    `json:"token"`
	}

	var user = user.User{}

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helpers.ErrorResponse{
			Success: false,
			Message: "validation error",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	// check db if username not found
	if err := database.DB.Where("username = ?", userRequest.Username).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Success: false,
			Message: "user not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Success: false,
			Message: "password invalid",
			Errors:  helpers.TranslateErrorMessage(err),
		})

		return
	}

	// generate token if succesfully logged in
	token := helpers.GenerateToken(user.Username)

	ctx.JSON(http.StatusOK, helpers.SuccessResponse{
		Success: true,
		Message: "logged in",
		Data: &userPayload{
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: utils.NowWIB(),
			UpdatedAt: utils.NowWIB(),
			Token:     token,
		},
	})

}
