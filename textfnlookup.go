package okay

import (
	"fmt"
	"net"
	"net/mail"
)

// TODO URL, URI, dates, credit cards, ...

type textConstraintFn func(*TextValue, constraint) (string, error)

var _TEXT_CONSTRAINT_FN_LOOKUP = map[string]textConstraintFn{
	"required": __required, "minlength": __minlength, "maxlength": __maxlength,
	"isemail": __isemail, "is": __is, "contains": __contains, "startswith": __startswith,
	"doesnotstartwith": __doesnotstartwith, "__endswith": __endswith, "doesnotendwith": __doesnotendwith,
	"isipv4": __isipv4, "isipv6": __isipv6, "isalpha": __isalpha, "isalphanumeric": __isalphanumeric,
	"isonlydigits": __isonlydigits,
}

func __required(v *TextValue, c constraint) (string, error) {
	if len(v.val) == 0 {
		return fmt.Sprintf("%s is required", v.fieldName), nil
	}
	return "", nil
}

func __minlength(v *TextValue, c constraint) (string, error) {
	return v.doMinLengthCheck(c)
}

func __maxlength(v *TextValue, c constraint) (string, error) {
	return v.doMaxLengthCheck(c)
}

func __isemail(v *TextValue, c constraint) (string, error) {
	_, err := mail.ParseAddress(v.val)
	if err != nil {
		return "invalid e-mail", nil
	}
	return "", nil
}

func __is(v *TextValue, c constraint) (string, error) {
	return v.doStringEquality(c)
}

func __contains(v *TextValue, c constraint) (string, error) {
	return v.doContainsCheck(c)
}

func __startswith(v *TextValue, c constraint) (string, error) {
	return v.doStartsWithCheck(c)
}

func __doesnotstartwith(v *TextValue, c constraint) (string, error) {
	return v.doDoesNotStartWithCheck(c)
}

func __endswith(v *TextValue, c constraint) (string, error) {
	return v.doEndsWithCheck(c)
}

func __doesnotendwith(v *TextValue, c constraint) (string, error) {
	return v.doDoesNotEndWithCheck(c)
}

func __isipv4(v *TextValue, c constraint) (string, error) {
	ip := net.ParseIP(v.val)
	if ip == nil {
		return fmt.Sprintf("%s must be a valid IPv4 address", v.fieldName), nil
	}
	return "", nil
}

func __isipv6(v *TextValue, c constraint) (string, error) {
	ip := net.ParseIP(v.val)
	if ip == nil {
		return fmt.Sprintf("%s must be a valid IPv6 address", v.fieldName), nil
	}
	return "", nil
}

func __isalpha(v *TextValue, c constraint) (string, error) {
	val := v.val
	for _, char := range val {
		ok := ((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z'))
		if !(ok) {
			return fmt.Sprintf("%s must not contain non-alpha characters ('%s')", v.fieldName, string(char)), nil
		}
	}
	return "", nil
}

func __isalphanumeric(v *TextValue, c constraint) (string, error) {
	val := v.val
	for _, char := range val {
		ok := ((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9'))
		if !(ok) {
			return fmt.Sprintf("%s must not contain non-alphanumeric characters ('%s')", v.fieldName, string(char)), nil
		}
	}
	return "", nil
}

func __isonlydigits(v *TextValue, c constraint) (string, error) {
	val := v.val
	for _, char := range val {
		ok := char >= '0' && char <= '9'
		if !(ok) {
			return fmt.Sprintf("%s must not contain non-numeric characters ('%s')", v.fieldName, string(char)), nil
		}
	}
	return "", nil
}
