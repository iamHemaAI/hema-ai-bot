package domain

import (
	"strconv"
	"strings"

	"github.com/iamHemaAI/hema-ai-bot/internal/tgbot/tgdto"
	"github.com/iamHemaAI/hema-ai-bot/pkg/validator"
)

type QuizState struct {
	Step int
	DTO  tgdto.BloodDTO
}

func parseAndSetInt(text string, dst *int) error {
	if err := validator.ValidateInt(text); err != nil {
		return err
	}
	*dst, _ = strconv.Atoi(text)
	return nil
}

func parseAndSetFloat(text string, dst *float64) error {
	if err := validator.ValidateFloat(text); err != nil {
		return err
	}
	*dst, _ = strconv.ParseFloat(strings.ReplaceAll(text, ",", "."), 64)
	return nil
}

var QuizSteps = []struct {
	Prompt string
	Parse  func(st *QuizState, text string) error
}{
	{
		"Введите гемоглобин (Hb), г/л (целое число):",
		func(st *QuizState, text string) error { return parseAndSetInt(text, &st.DTO.Hb) },
	},
	{
		"Введите эритроциты (RBC), 10¹²/л (число с точкой):",
		func(st *QuizState, text string) error { return parseAndSetFloat(text, &st.DTO.Rbc) },
	},
	{
		"Введите лейкоциты (WBC), 10⁹/л (число с точкой):",
		func(st *QuizState, text string) error { return parseAndSetFloat(text, &st.DTO.Wbc) },
	},
	{
		"Введите тромбоциты (PLT), 10⁹/л (целое число):",
		func(st *QuizState, text string) error { return parseAndSetInt(text, &st.DTO.Plt) },
	},
	{
		"Введите гематокрит (HCT), % (целое число):",
		func(st *QuizState, text string) error { return parseAndSetInt(text, &st.DTO.Hct) },
	},
	{
		"Введите средний объём эритроцита (MCV), фл (целое число):",
		func(st *QuizState, text string) error { return parseAndSetInt(text, &st.DTO.Mcv) },
	},
	{
		"Введите среднее содержание гемоглобина в эритроците (MCH), пг (целое число):",
		func(st *QuizState, text string) error { return parseAndSetInt(text, &st.DTO.Mch) },
	},
	{
		"Введите СОЭ (ESR), мм/ч (целое число):",
		func(st *QuizState, text string) error { return parseAndSetInt(text, &st.DTO.Esr) },
	},
	{
		"Введите глюкозу (Glucose), ммоль/л (число с точкой):",
		func(st *QuizState, text string) error { return parseAndSetFloat(text, &st.DTO.Glucose) },
	},
	{
		"Введите общий белок (Protein), г/л (целое число):",
		func(st *QuizState, text string) error { return parseAndSetInt(text, &st.DTO.Protein) },
	},
}
