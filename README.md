### Okay

Simple validation library for Go.

 ```go
type UserRegistrationPayload struct {
    Username, Email, Password, City string
    Age uint8
}

func (u UserRegistrationPayload) Okay() (ValidationErrors, error) {
    o := okay.New()

    o.Text(u.Username, "username").Required().MinLength(6)
    o.Text(u.Email, "email").Required().IsEmail().DoesNotEndWith(".edu")
	o.Text(u.City, "city").Required().Is("London")
    ...

    return o.Errors()
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var u UserRegistrationPayload
	json.NewEncoder(r.Body).Encode(&u)
	validationErrs, err := okay.Validate(u) // calls u.Okay, under the hood
	if err != nil {
		// unexpected error
		// should probably give an internal server error
	}
	if len(validationErrs) > 0 {
		// set headers 
		// send the errors
		// ...
	}
}
 ```