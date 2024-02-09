package models

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type InputStudent struct {
	FirstName      string        `json:"first_name" validate:"required"`
	LastName       string        `json:"last_name" validate:"required"`
	Age            int32         `json:"age" validate:"gte=0,lte=130"`
	Email          string        `json:"email" validate:"required,email"`
	Gender         string        `json:"gender" validate:"oneof=male female prefer_not_to"`
	FavouriteColor string        `json:"favouriteColor" validate:"iscolor"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      *InputAddress `json:"addresses" validate:"required,dive,required"`
}
type InputAddress struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var validate *validator.Validate

type InputUpdateStudent struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Age       int32  `json:"age" validate:"required"`
	Email     string `json:"email" validate:"required"`
}

func (s *InputUpdateStudent) Decode(data []byte) error {
	return json.Unmarshal(data, &s)
}
func (s *InputUpdateStudent) ValidateUpdateStudent(input InputUpdateStudent) error {
	validate = validator.New()
	err := validate.Struct(input)
	if err != nil {
		return err
	}
	return nil
}

func (s *InputStudent) ValidateStudent(input InputStudent) error {
	validate = validator.New()
	err := validate.Struct(input)
	if err != nil {
		return err
	}
	return nil
}
func (s *InputStudent) Decode(data []byte) error {
	return json.Unmarshal(data, &s)
}
