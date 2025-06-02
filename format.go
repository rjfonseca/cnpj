package cnpj

import "strings"

// Format formats a CNPJ string into the standard format: XX.XXX.XXX/XXXX-XX
// If the input string is not 14 characters long, it returns the string as is.
// It does not validate the content of the string, only its length.
func Format(c string) string {
	if len(c) != 14 {
		return c // Return as is if length is not 14
	}

	return c[:2] + "." + c[2:5] + "." + c[5:8] + "/" + c[8:12] + "-" + c[12:]
}

// Clean removes all non-digit characters from the input string.
// It retains only digits (0-9) and uppercase letters (A-Z).
// This is useful for cleaning up CNPJ strings before validation or formatting.
func Clean(c string) string {
	// Remove all non-digit characters
	sb := strings.Builder{}
	sb.Grow(14)

	for _, r := range c {
		if (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}
