package okay

/*
An interface for structs that are going to be validated using okay.Validate, to implement.
*/
type Validator interface {
	Okay() (ValidationErrors, error)
}

type O struct {
	texts []*TextValue
}

func New() *O {
	return &O{}
}

func (o *O) Errors() (ValidationErrors, error) {
	var res ValidationErrors
	for _, v := range o.texts {
		ex, err := v.validate()
		res = append(res, ex...)
		if err != nil {
			return res, err
		}
	}
	return res, nil
}

type ValidationErrors = []string

func Validate(input Validator) (ValidationErrors, error) {
	return input.Okay()
}

type constraint struct {
	name   string
	params []interface{}
}
