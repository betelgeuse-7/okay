### Okay

Simple validation library for Go.

 ```go

 type User struct {
    Username string 
    Email string
    Age uint
}

func (u User) Okay() []string {
	var errs ValidationErrors
	ex, err := okay.Text(u.username, "username").Required().MinLength(6).Errors()
	if err != nil {
		return errs, err
	}
	errs = append(errs, ex...)
	ex, err = okay.Text(u.email, "email").Required().IsEmail().Errors()
	errs = append(errs, ex...)
    ...
	return errs, err
}

u := User{
    Username: "user1",
    Email: "user1gmail",
    Age: 32,
}

errs := okay.Validate(u)

// []string{
//    "username length is 6 minimum" ,
//    "invalid email",
// }


 ```