package validation

import (
	"github.com/go-playground/validator"
)

type PayloadValidate interface {
	ValidateStruct(dto any) error
}
type PayloadValidation struct {
	validator *validator.Validate
}

func NewPayloadValidate() PayloadValidate {
	return &PayloadValidation{
		validator: validator.New(),
	}
}

func (pv *PayloadValidation) ValidateStruct(dto any) error {
	return pv.validator.Struct(dto)
}
