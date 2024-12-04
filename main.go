package main

import (
	"fmt"

	"github.com/danubiobwm/goEmailN/internal/domain/campaign"
	"github.com/go-playground/validator/v10"
)

func main() {
	campaign := campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(campaign)

	if err == nil {
		fmt.Println("Nenhum erro ")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			switch v.Tag() {
			case "required":
				println(v.StructField() + " is required: " + v.Tag())
			case "min":
				println(v.StructField() + " is required with min: " + v.Param())
			case "max":
				println(v.StructField() + " is required with max: " + v.Param())
			case "email":
				println(v.StructField() + " is invalid: ")
			}
			//println(v.StructField() + " is invalid: " + v.Tag())
		}
	}
}
