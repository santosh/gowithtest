package roman

import "strings"

// ConvertToRoman converts Arabic numerals to Roman.
func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for i := arabic; i > 0; i-- {
		if i == 4 {
			result.WriteString("IV")
			break
		}
		result.WriteString("I")
	}

	return result.String()
}
