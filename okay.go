package okay

/*
An interface for structs that are going to be validated using okay.Validate, to implement.
e.g.

type User struct {
    Username string
    Email string
    Age uint
}

func (u User) Okay() []string {
    var errs okay.ValidationErrors
    errs.Push(okay.Text(u.Username, okay.Required, okay.MinLength(6)))
    ...
    return errs
}
*/
type Validator interface {
	Okay() (ValidationErrors, error)
}

type ValidationErrors = []string

func Validate(input Validator) (ValidationErrors, error) {
	return input.Okay()
}

type constraint struct {
	name   string
	params []interface{}
}
