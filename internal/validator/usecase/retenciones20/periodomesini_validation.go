package retenciones20

import (
	"strconv"

	"github.com/gevs90/invoice-cfdi-validator/pkg/retenciones20"
	"github.com/gevs90/invoice-cfdi-validator/pkg/validator"
)

type MesIniUseCase struct {
	Retenciones retenciones20.Retenciones
}

func NewMesIniUseCase(retenciones retenciones20.Retenciones) *MesIniUseCase {
	return &MesIniUseCase{Retenciones: retenciones}
}

func (mi MesIniUseCase) MesIniValidation() validator.AttributeValidation {
	return validator.NewAttributeValidation(
		"Retenciones:Periodo:MesIni", mi.Retenciones.Periodo.MesIni,
		"Reten20122",
		"El campo MesIni no es menor o igual que el campo MesFin.",
		func(attributeValue string) (isValid bool) {
			mesIni, _ := strconv.Atoi(attributeValue)
			mesFin, _ := strconv.Atoi(mi.Retenciones.Periodo.MesFin)

			if mesFin >= mesIni {
				return true
			}
			return
		})
}
