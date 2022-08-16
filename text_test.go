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
