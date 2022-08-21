package okay

import "testing"

type user struct {
	username string
	email    string
}

func (u user) Okay() (ValidationErrors, error) {
	o := New()
	o.Text(u.username, "username").Required().IsAlphanumeric().MinLength(6)
	o.Text(u.email, "email").Required().IsEmail().DoesNotEndWith(".edu")
	return o.Errors()
}

func TestNewAPI(t *testing.T) {
	u := user{username: "he", email: "he@he.edu"}
	res, err := Validate(u)
	if err != nil {
		t.Logf("error: %s", err.Error())
		return
	}
	t.Logf("validation errors: %v\n", res)
}

type a struct {
	a string
	b string
}

func (a a) Okay() (ValidationErrors, error) {
	o := New()
	o.Text(a.a, "a").MinLength(6)
	o.Text(a.b, "b").MaxLength(12)
	return o.Errors()
}

func TestLengthCheck(t *testing.T) {
	a := a{a: "hfa", b: "1234567890123"}
	ex, err := Validate(a)
	if err != nil {
		t.Logf("error: %s\n", err.Error())
		return
	}
	t.Logf("Validation Errors: %v", ex)
}

func TestIsEmail(t *testing.T) {
	o := New()
	var email = "email@email.edu"
	o.Text(email, "email").IsEmail().DoesNotEndWith(".edu")
	ex, err := o.Errors()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(ex)
}

func TestIs(t *testing.T) {
	o := New()
	var x string = "non dirmi buena notte"
	o.Text(x, "x").Required().Is("ne dis pas que c'est la nuit")
	ex, err := o.Errors()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(ex)
}

func TestContains(t *testing.T) {
	o := New()
	var x string = "I will be there"
	o.Text(x, "x").Contains("obnoxious")
	ex, err := o.Errors()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(ex)
}

func TestStartsWith(t *testing.T) {
	o := New()
	var x string = "I will be there"
	o.Text(x, "x").StartsWith("cheerful")
	ex, err := o.Errors()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(ex)
}

func TestDoesNotStartWith(t *testing.T) {
	o := New()
	var x string = "I will be there"
	o.Text(x, "x").DoesNotStartWith("I wi")
	ex, err := o.Errors()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(ex)
}

func TestEndsWith(t *testing.T) {
	o := New()
	var x string = "I will be there"
	o.Text(x, "x").EndsWith("absolutely")
	ex, err := o.Errors()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(ex)
}

func TestDoesNotEndWith(t *testing.T) {
	o := New()
	var x string = "I will be there"
	o.Text(x, "x").DoesNotEndWith("there")
	ex, err := o.Errors()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(ex)
}
