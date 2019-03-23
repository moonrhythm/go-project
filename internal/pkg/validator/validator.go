package validator

import (
	"fmt"
	"unicode/utf8"

	"github.com/asaskevich/govalidator"
	"github.com/moonrhythm/validator"
)

// Validator type
type Validator struct {
	validator.Validator
}

type Error = validator.Error

// New creates new validator
func New() *Validator {
	return new(Validator)
}

// IsNotEmpty validates is the string empty
func (v *Validator) IsNotEmpty(field string, x string) {
	v.Must(x != "", fmt.Sprintf("'%s' must not empty", field))
}

// IsUUIDv4 validates is the string uuid v4
func (v *Validator) IsUUIDv4(field string, x string) {
	v.Must(govalidator.IsUUIDv4(x), fmt.Sprintf("'%s' must be uuid", field))
}

// MinLength validates minimum length of the string
func (v *Validator) MinLength(field string, min int, x string) {
	v.Must(utf8.RuneCountInString(x) >= min, fmt.Sprintf("'%s' must have at least %d characters", field, min))
}

// MaxLength validates maximum length of the string
func (v *Validator) MaxLength(field string, max int, x string) {
	v.Must(utf8.RuneCountInString(x) <= max, fmt.Sprintf("'%s' must have no more than %d characters", field, max))
}

// Min validates minimum value of int
func (v *Validator) Min(field string, min int, x int) {
	v.Must(x >= min, fmt.Sprintf("'%s' minimum value is %d", field, min))
}

// Max validates maximum value of int
func (v *Validator) Max(field string, max int, x int) {
	v.Must(x <= max, fmt.Sprintf("'%s' maximum value is %d", field, max))
}
