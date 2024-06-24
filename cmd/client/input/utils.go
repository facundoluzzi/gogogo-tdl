package input

import "unicode"

func splitFields(input string) []string {
	var fields []string
	var currentField []rune
	inQuotes := false

	for i, r := range input {
		switch {
		case r == '"':
			if inQuotes && (i == 0 || input[i-1] != '\\') {
				inQuotes = false
			} else {
				inQuotes = true
			}
		case unicode.IsSpace(r) && !inQuotes:
			if len(currentField) > 0 {
				fields = append(fields, string(currentField))
				currentField = nil
			}
		default:
			currentField = append(currentField, r)
		}
	}
	if len(currentField) > 0 {
		fields = append(fields, string(currentField))
	}
	return fields
}
