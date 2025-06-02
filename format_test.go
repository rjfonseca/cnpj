package cnpj

import "testing"

func TestFormat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "valid CNPJ",
			input:    "11222333000181",
			expected: "11.222.333/0001-81",
		},
		{
			name:     "valid CNPJ with letters",
			input:    "12ABC34501DE35",
			expected: "12.ABC.345/01DE-35",
		},
		{
			name:     "short input",
			input:    "1122233300018",
			expected: "1122233300018",
		},
		{
			name:     "long input",
			input:    "112223330001811",
			expected: "112223330001811",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "already formatted",
			input:    "11.222.333/0001-81",
			expected: "11.222.333/0001-81",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.input); got != tt.expected {
				t.Errorf("Format() = %v, want %v", got, tt.expected)
			}
		})
	}
}
func TestClean(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "already clean CNPJ",
			input:    "11222333000181",
			expected: "11222333000181",
		},
		{
			name:     "formatted CNPJ",
			input:    "11.222.333/0001-81",
			expected: "11222333000181",
		},
		{
			name:     "CNPJ with spaces",
			input:    "11 222 333 0001 81",
			expected: "11222333000181",
		},
		{
			name:     "CNPJ with special characters",
			input:    "11@222#333$0001*81",
			expected: "11222333000181",
		},
		{
			name:     "CNPJ with letters",
			input:    "11ABC33DE0001FG",
			expected: "11ABC33DE0001FG",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only special characters",
			input:    "@#$%^&*()_+-=",
			expected: "",
		},
		{
			name:     "lowercase letters should be removed",
			input:    "11abc33de0001fg",
			expected: "11330001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clean(tt.input); got != tt.expected {
				t.Errorf("Clean() = %v, want %v", got, tt.expected)
			}
		})
	}
}
