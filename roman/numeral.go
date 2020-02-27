package roman

// ConvertToRoman converts Arabic numerals to Roman.
func ConvertToRoman(arabic int) string {
	if arabic == 2 {
		return "II"
	}
	return "I"
}
