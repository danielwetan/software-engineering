package pkg

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidatorType struct {
	validate *validator.Validate
}

type ValidationMapData map[string]interface{}
type ValidationMapRule map[string]interface{}

// use a single instance of Validate, it caches struct info
var Validator *ValidatorType

func init() {
	v := validator.New(validator.WithRequiredStructEnabled())

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	Validator = &ValidatorType{validate: v}
}

// TODO:
// working on custom error message
// https://github.com/go-playground/validator/issues/559
func (v *ValidatorType) something() string {
	return "something"
}

// TODO:
// exctract error messages
// https://github.com/go-playground/validator/blob/master/_examples/simple/main.go
func (v *ValidatorType) ValidateStruct(s interface{}) error {
	err := v.validate.Struct(s)
	return err
}

func (v *ValidatorType) ValidateVar(field interface{}, tag string) error {
	err := v.validate.Var(field, tag)
	return err
}

// TODO:
// handle map error
// https://github.com/go-playground/validator/blob/master/_examples/map-validation/main.go
func (v *ValidatorType) ValidateMap(data, rules map[string]interface{}) error {
	err := v.validate.ValidateMap(data, rules)
	if len(err) > 0 {
		return errors.New("invalid parameter")
	}
	return nil
}
