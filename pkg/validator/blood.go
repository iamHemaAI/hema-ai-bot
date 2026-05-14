package validator

import (
	"fmt"
	"strings"
)

func ValidateBlood(
	valHb, valPlt, valHct, valMcv, valMch, valEsr, valProtein int,
	valGlucose, valRbc, valWbc float64,
) error {
	type intRange struct {
		name     string
		val      int
		min, max int
	}
	type floatRange struct {
		name     string
		val      float64
		min, max float64
	}

	intFields := []intRange{
		{"Гемоглобин (Hb)", valHb, 0, 300},
		{"Тромбоциты (PLT)", valPlt, 0, 3000},
		{"Гематокрит (HCT)", valHct, 0, 100},
		{"MCV", valMcv, 0, 200},
		{"MCH", valMch, 0, 100},
		{"СОЭ", valEsr, 0, 200},
		{"Общий белок", valProtein, 0, 200},
	}

	floatFields := []floatRange{
		{"Глюкоза", valGlucose, 0, 100},
		{"Эритроциты (RBC)", valRbc, 0, 10},
		{"Лейкоциты (WBC)", valWbc, 0, 200},
	}

	var errs []string

	for _, f := range intFields {
		if f.val < f.min || f.val > f.max {
			errs = append(errs, fmt.Sprintf("%s: значение %d вне допустимого диапазона [%d, %d]", f.name, f.val, f.min, f.max))
		}
	}

	for _, f := range floatFields {
		if f.val < f.min || f.val > f.max {
			errs = append(errs, fmt.Sprintf("%s: значение %.2f вне допустимого диапазона [%.2f, %.2f]", f.name, f.val, f.min, f.max))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%s", strings.Join(errs, "\n"))
	}

	return nil
}
