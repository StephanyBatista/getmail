package util

import (
	"fmt"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

//ValidateData valids a model throught its tag
func ValidateData(dataSet interface{}) error {
	err := validate.Struct(dataSet)

	if err != nil {

		//Validation syntax is invalid
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		// numberOfErros := len(err.(validator.ValidationErrors))
		// errors := make([]error, numberOfErros)

		for _, err := range err.(validator.ValidationErrors) {

			switch err.Tag() {
			case "required":
				//errors = append(errors, fmt.Errorf("The "+err.Field()+" is required"))
				return fmt.Errorf("The " + err.Field() + " is required")
			case "email":
				//errors = append(errors, fmt.Errorf("The "+err.Field()+" should be a valid email"))
				return fmt.Errorf("The " + err.Field() + " should be a valid email")
			case "eqfield":
				//errors = append(errors, fmt.Errorf("The "+err.Field()+" should be equal to the "+err.Param()))
				return fmt.Errorf("The " + err.Field() + " should be equal to the " + err.Param())
			default:
				//errors = append(errors, fmt.Errorf("The "+err.Field()+" is invalid"))
				return fmt.Errorf("The " + err.Field() + " is invalid")
			}
		}
	}
	return nil
}
