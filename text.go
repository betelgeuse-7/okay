package okay

import (
	"errors"
	"fmt"
	"strings"
)

type TextValue struct {
	val         string
	fieldName   string
	constraints []constraint
}

func (v *TextValue) newConstraint(name string, params ...interface{}) {
	c := constraint{name: name}
	if len(params) > 0 {
		c.params = params
	}
	v.constraints = append(v.constraints, c)
}

// a validator function to validate inputs of type string (aka text)
func Text(val string, fieldName string) *TextValue {
	v := &TextValue{val: val, fieldName: fieldName}
	return v
}

func (v *TextValue) Errors() (ValidationErrors, error) {
	return v.validate()
}

// Value is required. The value cannot have a zero value.
func (v *TextValue) Required() *TextValue {
	v.newConstraint("required")
	return v
}

// Set minimum length value for value.
func (v *TextValue) MinLength(length uint) *TextValue {
	v.newConstraint("minlength", length)
	return v
}

func (v *TextValue) MaxLength(length uint) *TextValue {
	v.newConstraint("maxlength", length)
	return v
}

func (v *TextValue) IsEmail() *TextValue {
	v.newConstraint("isemail")
	return v
}

// text value is "compare". (string equality)
func (v *TextValue) Is(compare string) *TextValue {
	v.newConstraint("is", compare)
	return v
}

func (v *TextValue) Contains(str string) *TextValue {
	v.newConstraint("contains", str)
	return v
}

func (v *TextValue) StartsWith(str string) *TextValue {
	v.newConstraint("startswith", str)
	return v
}

func (v *TextValue) DoesNotStartWith(str string) *TextValue {
	v.newConstraint("doesnotstartwith", str)
	return v
}

func (v *TextValue) EndsWith(str string) *TextValue {
	v.newConstraint("endswith", str)
	return v
}

func (v *TextValue) DoesNotEndWith(str string) *TextValue {
	v.newConstraint("doesnotendwith", str)
	return v
}

func (v *TextValue) IsIPv4() *TextValue {
	v.newConstraint("isipv4")
	return v
}

func (v *TextValue) IsIPv6() *TextValue {
	v.newConstraint("isipv6")
	return v
}

func (v *TextValue) IsAlpha() *TextValue {
	v.newConstraint("isalpha")
	return v
}

func (v *TextValue) IsAlphanumeric() *TextValue {
	v.newConstraint("isalphanumeric")
	return v
}

func (v *TextValue) validate() (ValidationErrors, error) {
	var errs ValidationErrors
	for _, c := range v.constraints {
		fn, ok := _TEXT_CONSTRAINT_FN_LOOKUP[c.name]
		if !(ok) {
			return errs, errors.New("unknown constraint: '" + c.name + "'")
		}
		validationErr, err := fn(v, c)
		if err != nil {
			return errs, err
		}
		if validationErr != "" {
			errs = append(errs, validationErr)
		}
	}
	return errs, nil
}

// TODO DRY

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

func (v *TextValue) doMaxLengthCheck(c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("doMaxLengthCheck: error: maximum length parameter is not given")
	}
	max := c.params[0].(uint)
	if uint(len(v.val)) > max {
		return fmt.Sprintf("maximum length for %s is %d", v.fieldName, max), nil
	}
	return "", nil
}

func (v *TextValue) doStringEquality(c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("doStringEquality: error: compare string is not given")
	}
	compare := c.params[0].(string)
	if v.val != compare {
		return fmt.Sprintf("%s must be '%s'", v.fieldName, compare), nil
	}
	return "", nil
}

func (v *TextValue) doContainsCheck(c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("doContainsCheck: error: substring is not given")
	}
	str := c.params[0].(string)
	ok := strings.Contains(v.val, str)
	if !(ok) {
		return fmt.Sprintf("%s must contain '%s'", v.fieldName, str), nil
	}
	return "", nil
}

func (v *TextValue) doStartsWithCheck(c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("doStartsWithCheck: error: prefix is not given")
	}
	str := c.params[0].(string)
	ok := strings.HasPrefix(v.val, str)
	if !(ok) {
		return fmt.Sprintf("%s must start with '%s'", v.fieldName, str), nil
	}
	return "", nil
}

func (v *TextValue) doDoesNotStartWithCheck(c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("doDoesNotStartWithCheck: error: prefix is not given")
	}
	str := c.params[0].(string)
	ok := strings.HasPrefix(v.val, str)
	if ok {
		return fmt.Sprintf("%s must not start with '%s'", v.fieldName, str), nil
	}
	return "", nil
}

func (v *TextValue) doEndsWithCheck(c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("doEndsWithCheck: error: suffix is not given")
	}
	str := c.params[0].(string)
	ok := strings.HasSuffix(v.val, str)
	if !(ok) {
		return fmt.Sprintf("%s must end with '%s'", v.fieldName, str), nil
	}
	return "", nil
}

func (v *TextValue) doDoesNotEndWithCheck(c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("doDoesNotEndWithCheck: error: suffix is not given")
	}
	str := c.params[0].(string)
	ok := strings.HasSuffix(v.val, str)
	if ok {
		return fmt.Sprintf("%s must not end with '%s'", v.fieldName, str), nil
	}
	return "", nil
}
