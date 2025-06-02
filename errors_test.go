package cnpj

import (
	"testing"
)

func TestErrInvalidCharacter_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      ErrInvalidCharacter
		expected string
	}{
		{
			name: "invalid character",
			err: ErrInvalidCharacter{
				Position:  5,
				Character: '@',
			},
			expected: "invalid character at position 5: @",
		},
		{
			name: "invalid character with special char",
			err: ErrInvalidCharacter{
				Position:  0,
				Character: '#',
			},
			expected: "invalid character at position 0: #",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.expected {
				t.Errorf("ErrInvalidCharacter.Error() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestErrUnexpectedDigit_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      ErrUnexpectedDigit
		expected string
	}{
		{
			name: "unexpected digit",
			err: ErrUnexpectedDigit{
				Position: 12,
				Expected: '5',
				Actual:   '7',
			},
			expected: "unexpected digit at position 12: expected 5, got 7",
		},
		{
			name: "unexpected digit at first position",
			err: ErrUnexpectedDigit{
				Position: 13,
				Expected: '1',
				Actual:   '2',
			},
			expected: "unexpected digit at position 13: expected 1, got 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.expected {
				t.Errorf("ErrUnexpectedDigit.Error() = %v, want %v", got, tt.expected)
			}
		})
	}
}
