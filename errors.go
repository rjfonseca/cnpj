package cnpj

import (
	"errors"
	"strconv"
)

var (
	ErrInvalidLength = errors.New("invalid CNPJ length, must be 14 characters")
)

// ErrInvalidCharacter represents an error for an invalid character in the CNPJ string.
// It includes the position of the character and the character itself.
// This error is used when validating the CNPJ and a character is found that is not a digit (0-9) or uppercase letter (A-Z).
type ErrInvalidCharacter struct {
	Position  int
	Character rune
}

func (e ErrInvalidCharacter) Error() string {
	return "invalid character at position " + strconv.Itoa(e.Position) + ": " + string(e.Character)
}

func newErrInvalidCharacter(pos int, r rune) *ErrInvalidCharacter {
	return &ErrInvalidCharacter{
		Position:  pos,
		Character: r,
	}
}

// ErrUnexpectedDigit represents an error for an unexpected digit in the CNPJ string.
// It indicates that the digit at a specific position does not match the expected value.
// This error is used when validating the CNPJ and the calculated digit does not match the provided one.
type ErrUnexpectedDigit struct {
	Position int
	Expected rune
	Actual   rune
}

func (e ErrUnexpectedDigit) Error() string {
	return "unexpected digit at position " + strconv.Itoa(e.Position) + ": expected " + string(e.Expected) + ", got " + string(e.Actual)
}

func newErrUnexpectedDigit(pos int, expected, actual rune) *ErrUnexpectedDigit {
	return &ErrUnexpectedDigit{
		Position: pos,
		Expected: expected,
		Actual:   actual,
	}
}
