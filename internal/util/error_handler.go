package util

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// ErrorResponse is a generic structure for API error responses
type ErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

// HandleValidationError processes validation errors and returns a structured error response
func HandleValidationError(err error) *ErrorResponse {
	var resp ErrorResponse
	resp.Errors = make(map[string]string)

	// Check if the error can be type asserted to validator.ValidationErrors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			fieldName := e.Field()
			tag := e.Tag()

			switch fieldName {
			case "Username":
				resp.Errors["username"] = "Username is required"
			case "Email":
				resp.Errors["email"] = "Valid email is required"
			case "Password":
				switch tag {
				case "required":
					resp.Errors["password"] = "Password is required"
				case "min":
					resp.Errors["password"] = "Password must be at least 8 characters"
				}
			case "ConfirmPassword":
				resp.Errors["confirm_password"] = "Password and confirm password must match"
			}
			// Extend the switch case for other fields and their potential validation tags as necessary
		}
	} else {
		// This is not a validation error from validator.v10
		resp.Errors["error"] = err.Error()
	}

	return &resp
}

// RespondWithError sends the error response with the given status code
func RespondWithError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, HandleValidationError(err))
}
