package helpers

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// If there is no value from .env given, assign the default value manually
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Hash the password to save in DB
func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

// Generate token
var jwtKey = []byte(GetEnv("JWT_SECRET", "SECRET_KEY"))

func GenerateToken(uname string) string {
	// Handle token expired time, set to 1 hour from now
	expTime := time.Now().Add(60 * time.Minute)

	// Create claim JWT
	claims := &jwt.RegisteredClaims{
		Subject:   uname,
		ExpiresAt: jwt.NewNumericDate(expTime),
	}

	fmt.Println("claims : ", claims)
	// create a new token with the claims created
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	fmt.Println("token : ", token)

	return token
}

// JSON Responses
type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}

// Error Translators
func TranslateErrorMessage(err error) map[string]string {

	errorsMap := make(map[string]string)

	// Handle validations
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			field := fieldError.Field() // save error field
			switch fieldError.Tag() {
			case "required":
				errorsMap[field] = fmt.Sprintf("%s is required", field) // error if field is none
			case "email":
				errorsMap[field] = "Invalid email format" // email not valid
			case "unique":
				errorsMap[field] = fmt.Sprintf("%s already exists", field) // data already exists
			case "min":
				errorsMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param()) // data too short
			case "max":
				errorsMap[field] = fmt.Sprintf("%s must be at most %s characters", field, fieldError.Param()) // data too long
			case "numeric":
				errorsMap[field] = fmt.Sprintf("%s must be a number", field) // data must be a number
			default:
				errorsMap[field] = "Invalid value" // for another invalid value
			}
		}
	}

	// Handle GORM error for duplicate entry
	if err != nil {

		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				errorsMap["Username"] = "Username already exists"
			}
			if strings.Contains(err.Error(), "email") {
				errorsMap["Email"] = "Email already exists"
			}
		} else if err == gorm.ErrRecordNotFound {
			// if data not found in DB
			errorsMap["Error"] = "Record not found"
		}
	}

	return errorsMap
}

// IsDuplicateEntryError detect whether error from DB is duplicate entry
func IsDuplicateEntryError(err error) bool {
	// checking whether error is duplicate entry
	return err != nil && strings.Contains(err.Error(), "Duplicate entry")
}
