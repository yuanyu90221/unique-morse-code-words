package unique_morse_representation

var MORSE_CODE = [...]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	morseMap := make(map[string]int)
	for _, word := range words {
		morseMap[translateStrToMorse(word)] += 1
	}
	return len(morseMap)
}

func translateStrToMorse(word string) string {
	result := ""
	for _, r := range word {
		result += MORSE_CODE[r-'a']
	}
	return result
}
