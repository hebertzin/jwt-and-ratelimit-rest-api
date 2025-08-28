package validation

import (
	"github.com/go-playground/validator"
)

type PayloadValidation struct{}

func (pv *PayloadValidation) ValidateStruct(dto any) error {
	return validator.New().Struct(dto)
}
