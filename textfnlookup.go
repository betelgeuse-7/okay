package okay

import (
	"errors"
	"fmt"
	"net"
	"net/mail"
	"strings"
)

// TODO URL, URI, dates, credit cards, ...

type textConstraintFn func(*TextValue, constraint) (string, error)

var _TEXT_CONSTRAINT_FN_LOOKUP = map[string]textConstraintFn{
	"required": required, "minlength": minLength, "maxlength": maxLength,
	"isemail": isEmail, "is": is, "contains": contains, "startswith": startsWith,
	"doesnotstartwith": doesNotStartWith, "endswith": endsWith, "doesnotendwith": doesNotEndWith,
	"isipv4": isIPv4, "isipv6": isIPv6, "isalpha": isAlpha, "isalphanumeric": isAlphanumeric,
	"isonlydigits": isOnlyDigits,
}

type lengthCheckOption uint8

const (
	_MIN lengthCheckOption = iota
	_MAX
)

func (v *TextValue) doLengthCheck(c constraint, opt lengthCheckOption) (string, error) {
	if len(c.params) == 0 {
		fnName := "Max"
		if opt == _MIN {
			fnName = "Min"
		}
		return "", fmt.Errorf("do%sLengthCheck: error: length parameter is not given", fnName)
	}
	number := c.params[0].(uint)
	expr := uint(len(v.val)) < number
	if opt == _MIN {
		if expr {
			return fmt.Sprintf("minimum length for %s is %d", v.fieldName, number), nil
		}
	} else {
		if !(expr) {
			return fmt.Sprintf("maximum length for %s is %d", v.fieldName, number), nil
		}
	}
	return "", nil
}

func required(v *TextValue, c constraint) (string, error) {
	if len(v.val) == 0 {
		return fmt.Sprintf("%s is required", v.fieldName), nil
	}
	return "", nil
}

func minLength(v *TextValue, c constraint) (string, error) {
	return v.doLengthCheck(c, _MIN)
}

func maxLength(v *TextValue, c constraint) (string, error) {
	return v.doLengthCheck(c, _MAX)
}

func isEmail(v *TextValue, c constraint) (string, error) {
	_, err := mail.ParseAddress(v.val)
	if err != nil {
		return "invalid e-mail", nil
	}
	return "", nil
}

func is(v *TextValue, c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("doStringEquality: error: compare string is not given")
	}
	compare := c.params[0].(string)
	if v.val != compare {
		return fmt.Sprintf("%s must be '%s'", v.fieldName, compare), nil
	}
	return "", nil
}

func contains(v *TextValue, c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("contains: error: substring is not given")
	}
	str := c.params[0].(string)
	ok := strings.Contains(v.val, str)
	if !(ok) {
		return fmt.Sprintf("%s must contain '%s'", v.fieldName, str), nil
	}
	return "", nil
}

func startsWith(v *TextValue, c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("startsWith: error: prefix is not given")
	}
	str := c.params[0].(string)
	ok := strings.HasPrefix(v.val, str)
	if !(ok) {
		return fmt.Sprintf("%s must start with '%s'", v.fieldName, str), nil
	}
	return "", nil
}

func doesNotStartWith(v *TextValue, c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("doesNotStartWith: error: prefix is not given")
	}
	str := c.params[0].(string)
	ok := strings.HasPrefix(v.val, str)
	if ok {
		return fmt.Sprintf("%s must not start with '%s'", v.fieldName, str), nil
	}
	return "", nil
}

func endsWith(v *TextValue, c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("endsWith: error: suffix is not given")
	}
	str := c.params[0].(string)
	ok := strings.HasSuffix(v.val, str)
	if !(ok) {
		return fmt.Sprintf("%s must end with '%s'", v.fieldName, str), nil
	}
	return "", nil
}

func doesNotEndWith(v *TextValue, c constraint) (string, error) {
	if len(c.params) == 0 {
		return "", errors.New("doesNotEndWith: error: suffix is not given")
	}
	str := c.params[0].(string)
	ok := strings.HasSuffix(v.val, str)
	if ok {
		return fmt.Sprintf("%s must not end with '%s'", v.fieldName, str), nil
	}
	return "", nil
}

func isIPv4(v *TextValue, c constraint) (string, error) {
	ip := net.ParseIP(v.val)
	if ip == nil {
		return fmt.Sprintf("%s must be a valid IPv4 address", v.fieldName), nil
	}
	return "", nil
}

func isIPv6(v *TextValue, c constraint) (string, error) {
	ip := net.ParseIP(v.val)
	if ip == nil {
		return fmt.Sprintf("%s must be a valid IPv6 address", v.fieldName), nil
	}
	return "", nil
}

func isAlpha(v *TextValue, c constraint) (string, error) {
	val := v.val
	for _, char := range val {
		ok := ((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z'))
		if !(ok) {
			return fmt.Sprintf("%s must not contain non-alpha characters ('%s')", v.fieldName, string(char)), nil
		}
	}
	return "", nil
}

func isAlphanumeric(v *TextValue, c constraint) (string, error) {
	val := v.val
	for _, char := range val {
		ok := ((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9'))
		if !(ok) {
			return fmt.Sprintf("%s must not contain non-alphanumeric characters ('%s')", v.fieldName, string(char)), nil
		}
	}
	return "", nil
}

func isOnlyDigits(v *TextValue, c constraint) (string, error) {
	val := v.val
	for _, char := range val {
		ok := char >= '0' && char <= '9'
		if !(ok) {
			return fmt.Sprintf("%s must not contain non-numeric characters ('%s')", v.fieldName, string(char)), nil
		}
	}
	return "", nil
}
