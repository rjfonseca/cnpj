package cnpj

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		cnpj      string
		expectErr bool
	}{
		{
			name: "valid CNPJ",
			cnpj: "11222333000181",
		},
		{
			name: "valid alphanumeric CNPJ",
			cnpj: "12ABC34501DE35",
		},
		{
			name: "valid CNPJ with zero in DV",
			cnpj: "84877190000110",
		},
		{
			name:      "invalid length",
			cnpj:      "1122233300018",
			expectErr: true,
		},
		{
			name:      "invalid first digit",
			cnpj:      "11222333000191",
			expectErr: true,
		},
		{
			name:      "invalid second digit",
			cnpj:      "11222333000189",
			expectErr: true,
		},
		{
			name:      "invalid characters",
			cnpj:      "a1222333000182",
			expectErr: true,
		},
		{
			name:      "invalid second digit characters",
			cnpj:      "1122233300018A",
			expectErr: true,
		},
		{
			name:      "invalid first digit characters",
			cnpj:      "112223330001A1",
			expectErr: true,
		},
		{
			name:      "empty string",
			cnpj:      "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.cnpj)
			if tt.expectErr && err == nil {
				t.Errorf("Validate() = nil, want error for %s", tt.cnpj)
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Validate() = %v, want nil for %s", err, tt.cnpj)
			}
		})
	}
}

func Test_toInt(t *testing.T) {
	tests := []struct {
		name     string
		r        rune
		expected int
	}{
		{
			name:     "valid digit",
			r:        '5',
			expected: 5,
		},
		{
			name:     "valid zero",
			r:        '0',
			expected: 0,
		},
		{
			name:     "character A",
			r:        'A',
			expected: 17,
		},
		{
			name:     "character Z",
			r:        'Z',
			expected: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toInt(tt.r); got != tt.expected {
				t.Errorf("toInt() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func Test_isValidDigit(t *testing.T) {
	tests := []struct {
		name     string
		r        rune
		expected bool
	}{
		{
			name:     "valid digit",
			r:        '5',
			expected: true,
		},
		{
			name:     "valid zero",
			r:        '0',
			expected: true,
		},
		{
			name:     "valid character A",
			r:        'A',
			expected: true,
		},
		{
			name:     "invalid character a",
			r:        'a',
			expected: false,
		},
		{
			name:     "invalid character",
			r:        '!',
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidDigit(tt.r); got != tt.expected {
				t.Errorf("isValidDigit() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		cnpj     string
		expected bool
	}{
		{
			name:     "valid CNPJ",
			cnpj:     "11222333000181",
			expected: true,
		},
		{
			name:     "valid alphanumeric CNPJ",
			cnpj:     "12ABC34501DE35",
			expected: true,
		},
		{
			name:     "valid CNPJ with zero in DV",
			cnpj:     "84877190000110",
			expected: true,
		},
		{
			name:     "invalid length",
			cnpj:     "1122233300018",
			expected: false,
		},
		{
			name:     "invalid first digit",
			cnpj:     "11222333000191",
			expected: false,
		},
		{
			name:     "invalid second digit",
			cnpj:     "11222333000182",
			expected: false,
		},
		{
			name:     "invalid characters",
			cnpj:     "a1222333000182",
			expected: false,
		},
		{
			name:     "empty string",
			cnpj:     "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.cnpj); got != tt.expected {
				t.Errorf("IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Validate("12ABC34501DE35")
	}
}
