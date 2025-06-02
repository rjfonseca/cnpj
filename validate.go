package cnpj

import (
	"errors"
)

var (
	ErrInvalidLength = errors.New("invalid CNPJ length, must be 14 characters")
)

var (
	weights = [13]int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5, 6}
)

// IsValid checks if the provided CNPJ string is valid.
// It returns true if the CNPJ is valid, and false otherwise.
func IsValid(c string) bool {
	return Validate(c) == nil
}

// Validate checks if the provided CNPJ string is valid.
// It returns an error if the CNPJ is invalid, and nil if it is valid.
// The regular expression that defines a valid CNPJ is [A-Z0-9]{12}[0-9]{2}.
func Validate(str string) error {
	if len(str) != 14 {
		return ErrInvalidLength
	}

	cnpj := []rune(str)

	if cnpj[12] < '0' || cnpj[12] > '9' {
		return newErrInvalidCharacter(12, cnpj[12])
	}
	if cnpj[13] < '0' || cnpj[13] > '9' {
		return newErrInvalidCharacter(13, cnpj[13])
	}

	// Validate first digit
	expectedDigit, err := calculateDigit(cnpj[:12])
	if err != nil {
		return err
	}
	if cnpj[12] != expectedDigit {
		return newErrUnexpectedDigit(12, expectedDigit, cnpj[12])
	}

	// Validate second digit
	// The error is ignored here because we already validated the first 12
	// characters and the first verification digit
	expectedDigit, _ = calculateDigit(cnpj[:13])

	if cnpj[13] != expectedDigit {
		return newErrUnexpectedDigit(13, expectedDigit, cnpj[13])
	}

	return nil
}

func calculateDigit(partialCNPJ []rune) (rune, error) {
	sum := 0
	for i, j := len(partialCNPJ)-1, 0; i >= 0; i, j = i-1, j+1 {
		if !isValidDigit(partialCNPJ[i]) {
			return '*', newErrInvalidCharacter(i, partialCNPJ[i])
		}
		sum += toInt(partialCNPJ[i]) * weights[j]
	}
	remainder := sum % 11
	if remainder < 2 {
		return '0', nil
	}
	return '0' + rune(11-remainder), nil
}

func isValidDigit(r rune) bool {
	return (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z')
}

func toInt(r rune) int {
	return int(r - '0')
}
