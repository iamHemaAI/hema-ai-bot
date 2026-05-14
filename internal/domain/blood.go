package domain

import (
	"github.com/google/uuid"
	bloodtypes "github.com/iamHemaAI/hema-ai-bot/internal/domain/blood_types"
	buityrutime "github.com/iamHemaAI/hema-ai-bot/internal/domain/buity_ru_time"
	"github.com/iamHemaAI/hema-ai-bot/pkg/apperrors"
	"github.com/iamHemaAI/hema-ai-bot/pkg/validator"
)

type Blood struct {
	ID     uuid.UUID
	UserID int64

	Hb      *bloodtypes.HB
	Rbc     *bloodtypes.RBC
	Wbc     *bloodtypes.WBC
	Plt     *bloodtypes.PLT
	Hct     *bloodtypes.HCT
	Mcv     *bloodtypes.MCV
	Mch     *bloodtypes.MCH
	Esr     *bloodtypes.ESR
	Glucose *bloodtypes.GLUCOSE
	Protein *bloodtypes.PROTEIN

	CreatedAt buityrutime.BuityRuTime
}

func CreateBlood(
	valUserID int64,
	valHb, valPlt, valHct, valMcv, valMch, valEsr, valProtein int,
	valGlucose, valRbc, valWbc float64,
) (*Blood, error) {
	err := validator.ValidateBlood(valHb, valPlt, valHct, valMcv, valMch, valEsr, valProtein, valGlucose, valRbc, valWbc)
	if err != nil {
		return nil, apperrors.NewDomainValidationError(err, err.Error())
	}

	return &Blood{
		ID:     uuid.New(),
		UserID: valUserID,

		Hb:      bloodtypes.NewHb(valHb),
		Rbc:     bloodtypes.NewRbc(valRbc),
		Wbc:     bloodtypes.NewWBC(valWbc),
		Plt:     bloodtypes.NewPLT(valPlt),
		Hct:     bloodtypes.NewHCT(valHct),
		Mcv:     bloodtypes.NewMCV(valMcv),
		Mch:     bloodtypes.NewMCH(valMch),
		Esr:     bloodtypes.NewESR(valEsr),
		Glucose: bloodtypes.NewGLUCOSE(valGlucose),
		Protein: bloodtypes.NewProtein(valProtein),

		CreatedAt: buityrutime.Now(),
	}, nil
}
