package valider

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var Valider *validator.Validate

func InitValider() {
	Valider = validator.New()
	Valider.RegisterValidation("username", validateUsername)
}

func validateUsername(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+(_YT[0-9]+)?$`)
	return re.MatchString(fl.Field().String())
}
