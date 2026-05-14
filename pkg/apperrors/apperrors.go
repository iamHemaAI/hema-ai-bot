package apperrors

import "errors"

// repository
var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

// domain

type DomainValidationError struct {
	err error
	msg string
}

var ErrValidationBlood = errors.New("ошибка валидации анализа крови")

func (e *DomainValidationError) Error() string {
	return e.msg
}

func (e *DomainValidationError) Unwrap() error {
	return e.err
}

func newDomainValidationError(err error, msg string) *DomainValidationError {
	return &DomainValidationError{
		err: err,
		msg: msg,
	}
}

func NewDomainValidationError(err error, msg string) *DomainValidationError {
	return newDomainValidationError(err, msg)
}

var (
	ErrInvalidHB      = newDomainValidationError(ErrValidationBlood, "HB не валиден")
	ErrInvalidRBC     = newDomainValidationError(ErrValidationBlood, "RBC не валиден")
	ErrInvalidWBC     = newDomainValidationError(ErrValidationBlood, "WBC не валиден")
	ErrInvalidPLT     = newDomainValidationError(ErrValidationBlood, "PLT не валиден")
	ErrInvalidHCT     = newDomainValidationError(ErrValidationBlood, "HCT не валиден")
	ErrInvalidMCV     = newDomainValidationError(ErrValidationBlood, "MCV не валиден")
	ErrInvalidMCH     = newDomainValidationError(ErrValidationBlood, "MCH не валиден")
	ErrInvalidESR     = newDomainValidationError(ErrValidationBlood, "ESR не валиден")
	ErrInvalidGLUCOSE = newDomainValidationError(ErrValidationBlood, "GLUCOSE не валиден")
	ErrInvalidPROTEIN = newDomainValidationError(ErrValidationBlood, "PROTEIN не валиден")
)
