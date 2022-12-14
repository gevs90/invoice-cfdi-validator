package validator_test

import (
	"fmt"
	"testing"

	"github.com/gevs90/invoice-cfdi-validator/pkg/validator"
	"github.com/stretchr/testify/assert"
)

func TestErrAttrValidation_All(t *testing.T) {
	errAttrValidation := validator.ErrAttrValidation{
		Name:       "Total",
		Value:      "100.00",
		Code:       "CFDI5000",
		ErrMessage: "Formato invalido.",
	}
	errText := fmt.Sprintf("%s - %s", errAttrValidation.GetCode(), errAttrValidation.GetMessage())

	assert.Equal(t, errAttrValidation.GetCode(), errAttrValidation.Code)
	assert.Equal(t, errAttrValidation.GetMessage(), errAttrValidation.ErrMessage)
	assert.Equal(t, errAttrValidation.Error(), errText)
}
