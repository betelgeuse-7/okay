package okay

import "testing"

type user struct {
	username string
	email    string
}

func (u user) Okay() (ValidationErrors, error) {
	var errs ValidationErrors
	ex, err := Text(u.username, "username").Required().MinLength(6).Errors()
	if err != nil {
		return errs, err
	}
	errs = append(errs, ex...)
	ex, err = Text(u.email, "email").Required().IsEmail().Errors()
	errs = append(errs, ex...)
	return errs, err
}

func TestText(t *testing.T) {
	u := user{username: "hey", email: "hey@gmailcom"}
	ex, err := Validate(u)
	if err != nil {
		t.Logf("error: %s", err.Error())
		return
	}
	t.Logf("validation errors: %v", ex)
}

type a struct {
	username, email, x string
}

func (a a) Okay() (ValidationErrors, error) {
	errs := ValidationErrors{}
	ex, err := Text(a.username, "username").Required().StartsWith("ali").
		DoesNotEndWith("z").MinLength(3).Errors()
	if err != nil {
		return ex, err
	}
	errs = append(errs, ex...)
	ex, err = Text(a.email, "email").Required().IsEmail().StartsWith("ali").
		DoesNotStartWith("zali").MaxLength(35).Errors()
	if err != nil {
		return ex, err
	}
	errs = append(errs, ex...)
	ex, err = Text(a.x, "x").Is("the quick brown fox jumped over the lazy dogg").Errors()
	if err != nil {
		return ex, err
	}
	errs = append(errs, ex...)
	return errs, nil
}

func TestText2(t *testing.T) {
	u := a{username: "alice123", email: "zalice@alice.com", x: "the quick brown fox jumped over the lazy dog"}
	errs, err := Validate(u)
	if err != nil {
		t.Logf("error: %s", err.Error())
		return
	}
	t.Logf("validation errors: %v", errs)
}

type b struct {
	ipv4, ipv6 string
}

func (b b) Okay() (ValidationErrors, error) {
	errs := ValidationErrors{}
	ex, err := Text(b.ipv4, "ipv4").Required().IsIPv4().Errors()
	if err != nil {
		return ex, err
	}
	errs = append(errs, ex...)
	ex, err = Text(b.ipv6, "ipv6").Required().IsIPv6().Errors()
	if err != nil {
		return ex, err
	}
	errs = append(errs, ex...)
	return errs, nil
}

func TestIP(t *testing.T) {
	b := b{ipv4: "19.117.63.1000000126", ipv6: "684D:1111:222:3333:4444:5555:6:77"}
	errs, err := Validate(b)
	if err != nil {
		t.Logf("error: %s", err.Error())
		return
	}
	t.Logf("validation errors: %v", errs)
}

type c struct {
	cc string
}

func (c c) Okay() (ValidationErrors, error) {
	ex, err := Text(c.cc, "cc").Required().IsAlphanumeric().Errors()
	if err != nil {
		return ex, err
	}
	return ex, nil
}

func TestAlphanumeric(t *testing.T) {
	c := c{cc: "hello😀"}
	ex, err := Validate(c)
	if err != nil {
		t.Logf("error: %s\n", err.Error())
		return
	}
	t.Logf("validation errors: %v\n", ex)
}

type d struct {
	dd string
}

func (d d) Okay() (ValidationErrors, error) {
	ex, err := Text(d.dd, "dd").Required().IsAlpha().Errors()
	if err != nil {
		return ex, err
	}
	return ex, err
}

func TestAlpha(t *testing.T) {
	d := d{dd: "hey90"}
	ex, err := Validate(d)
	if err != nil {
		t.Logf("error: %s\n", err.Error())
		return
	}
	t.Logf("validation errors: %v\n", ex)
}

type e struct {
	ee string
}

func (e e) Okay() (ValidationErrors, error) {
	ex, err := Text(e.ee, "ee").Required().IsOnlyDigits().Errors()
	if err != nil {
		return ex, err
	}
	return ex, err
}

func TestIsOnlyDigits(t *testing.T) {
	e := e{ee: "958013h68913"}
	ex, err := Validate(e)
	if err != nil {
		t.Logf("error: %s\n", err.Error())
		return
	}
	t.Logf("validation errors: %v\n", ex)
}
