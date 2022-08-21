package okay

import (
	"errors"
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
func (o *O) Text(val string, fieldName string) *TextValue {
	v := &TextValue{val: val, fieldName: fieldName}
	o.texts = append(o.texts, v)
	return v
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

func (v *TextValue) IsOnlyDigits() *TextValue {
	v.newConstraint("isonlydigits")
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
