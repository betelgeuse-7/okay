package okay

import (
	"errors"
	"fmt"
	"net/mail"
)

type TextValue struct {
	val         string
	fieldName   string
	constraints []constraint
}

// a validator function to validate inputs of type string (aka text)
func Text(val string, fieldName string) *TextValue {
	v := &TextValue{val: val, fieldName: fieldName}
	return v
}

// Value is required. The value cannot have a zero value.
func (v *TextValue) Required() *TextValue {
	v.constraints = append(v.constraints, constraint{
		name: "required",
	})
	return v
}

// Set minimum length value for value.
func (v *TextValue) MinLength(length uint) *TextValue {
	v.constraints = append(v.constraints, constraint{
		name:   "minlength",
		params: []interface{}{length},
	})
	return v
}

func (v *TextValue) IsEmail() *TextValue {
	v.constraints = append(v.constraints, constraint{
		name: "isemail",
	})
	return v
}

func (v *TextValue) Errors() (ValidationErrors, error) {
	return v.validate()
}

func (v *TextValue) validate() (ValidationErrors, error) {
	var errs ValidationErrors
	for _, c := range v.constraints {
		switch c.name {
		case "required":
			if len(v.val) == 0 {
				errs = append(errs, fmt.Sprintf("%s is required", v.fieldName))
			}
		case "minlength":
			validationErr, err := v.doMinLengthCheck(c)
			if err != nil {
				return errs, err
			}
			if validationErr != "" {
				errs = append(errs, validationErr)
			}
		case "isemail":
			_, err := mail.ParseAddress(v.val)
			if err != nil {
				errs = append(errs, "invalid email")
			}
		}
	}
	return errs, nil
}

func (v *TextValue) doMinLengthCheck(c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("doMinLengthCheck: error: minimum length parameter is not given")
	}
	min := c.params[0].(uint)
	if uint(len(v.val)) < min {
		return fmt.Sprintf("minimum length for %s is %d", v.fieldName, min), nil
	}
	return "", nil
}
