package mappers

import (
	"github.com/google/uuid"
	"github.com/iamHemaAI/hema-ai-bot/internal/domain"
	bloodtypes "github.com/iamHemaAI/hema-ai-bot/internal/domain/blood_types"
	buityrutime "github.com/iamHemaAI/hema-ai-bot/internal/domain/buity_ru_time"
	"github.com/iamHemaAI/hema-ai-bot/internal/domain/records"
	"github.com/iamHemaAI/hema-ai-bot/internal/tgbot/tgdto"
)

func BloodToRecord(blood *domain.Blood) *records.Blood {
	return &records.Blood{
		ID:     blood.ID.String(),
		UserID: blood.UserID,

		Hb:      blood.Hb.Value(),
		Rbc:     blood.Rbc.Value(),
		Wbc:     blood.Wbc.Value(),
		Plt:     blood.Plt.Value(),
		Hct:     blood.Hct.Value(),
		Mcv:     blood.Mcv.Value(),
		Mch:     blood.Mch.Value(),
		Esr:     blood.Esr.Value(),
		Glucose: blood.Glucose.Value(),
		Protein: blood.Protein.Value(),

		CreatedAt: string(blood.CreatedAt),
	}
}

func RecordToBlood(record *records.Blood) *domain.Blood {
	return &domain.Blood{
		ID:     uuid.MustParse(record.ID),
		UserID: record.UserID,

		Hb:      bloodtypes.NewHb(record.Hb),
		Rbc:     bloodtypes.NewRbc(record.Rbc),
		Wbc:     bloodtypes.NewWBC(record.Wbc),
		Plt:     bloodtypes.NewPLT(record.Plt),
		Hct:     bloodtypes.NewHCT(record.Hct),
		Mcv:     bloodtypes.NewMCV(record.Mcv),
		Mch:     bloodtypes.NewMCH(record.Mch),
		Esr:     bloodtypes.NewESR(record.Esr),
		Glucose: bloodtypes.NewGLUCOSE(record.Glucose),
		Protein: bloodtypes.NewProtein(record.Protein),

		CreatedAt: buityrutime.BuityRuTime(record.CreatedAt),
	}
}

func BloodDTOtoBlood(dto *tgdto.BloodDTO) (*domain.Blood, error) {
	b, err := domain.CreateBlood(
		dto.UserID,
		dto.Hb,
		dto.Plt,
		dto.Hct,
		dto.Mcv,
		dto.Mch,
		dto.Esr,
		dto.Protein,
		dto.Glucose,
		dto.Rbc,
		dto.Wbc,
	)
	if err != nil {
		return nil, err
	}

	return b, nil
}
