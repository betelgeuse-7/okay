```go

type UserRegistrationPayload struct {
    Username, Email, Password, City string
    Age uint8
}

func (u UserRegistrationPayload) Okay() (ValidationErrors, error) {
    o := okay.New()

    o.Text(u.Username, "username").Required().MinLength(6)
    o.Text(u.Email, "email").Required().IsEmail().DoesNotEndWith(".edu")
    ...

    return o.Errors()
}



```