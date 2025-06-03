package cnpj

import "math/rand"

// Generate creates a random CNPJ number.
func Generate() string {
	// This function generates a random CNPJ number.
	// The CNPJ is a 14-character string, where the first 12 characters are digits,
	// and the last two are verification digits calculated from the first 12.
	cnpj := make([]rune, 14)

	// Generate the first 12 digits randomly
	for i := 0; i < 12; i++ {
		if rand.Intn(2) == 0 {
			cnpj[i] = '0' + rune(rand.Intn(10)) // Random digit from '0' to '9'
		} else {
			cnpj[i] = 'A' + rune(rand.Intn(26)) // Random uppercase letter from 'A' to 'Z'
		}
	}

	cnpj[12] = calculateDigitUnsafe(cnpj[:12])
	cnpj[13] = calculateDigitUnsafe(cnpj[:13])

	return string(cnpj)
}
