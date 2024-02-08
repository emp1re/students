package models

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type Student struct {
	FirstName      string   `json:"first_name" validate:"required"`
	LastName       string   `json:"last_name" validate:"required"`
	Age            uint8    `json:"age" validate:"gte=0,lte=130"`
	Email          string   `json:"email" validate:"required,email"`
	Gender         string   `json:"gender" validate:"oneof=male female prefer_not_to"`
	FavouriteColor string   `json:"favouriteColor" validate:"iscolor"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      *Address `json:"addresses" validate:"required,dive,required"`
}
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var validate *validator.Validate

func (s *Student) ValidateStudent(input Student) error {
	validate = validator.New()
	err := validate.Struct(input)
	if err != nil {
		return err
	}
	return nil
}
func (s *Student) Decode(data []byte) error {
	return json.Unmarshal(data, &s)
}
