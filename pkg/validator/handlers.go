package validator

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateFloat(text string) error {
	if text == "" {
		return fmt.Errorf("Значение не может быть пустым")
	}
	normalized := strings.ReplaceAll(text, ",", ".")

	if !strings.Contains(normalized, ".") {
		return fmt.Errorf("Ожидается дробное число (например: 3.14), получено: %q", text)
	}
	_, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("Ожидается дробное число (например: 3.14), получено: %q", text)
	}
	return nil
}

func ValidateInt(text string) error {
	if text == "" {
		return fmt.Errorf("Значение не может быть пустым")
	}
	_, err := strconv.Atoi(text)
	if err != nil {
		return fmt.Errorf("Ожидается целое число, получено: %q", text)
	}
	return nil
}
