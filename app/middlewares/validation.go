package middlewares

import (
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func InitValidator() *validator.Validate {
	validate := validator.New()

	validate.RegisterValidation("containsNumber", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`[0-9]`, fl.Field().String())
		return match
	})

	validate.RegisterValidation("containsSpecialCharacter", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`[!@#$%^&*()]`, fl.Field().String())
		return match
	})

	validate.RegisterValidation("validDate", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-(\d{4})$`, fl.Field().String())
		return match
	})

	return validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
