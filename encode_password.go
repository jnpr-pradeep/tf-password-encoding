package main

import (
	"fmt"
	"strings"
)

// EncodeSpecialCharsForTerraform encodes special characters according to the provided mapping.
func EncodeSpecialCharsForTerraform(input string) string {
	// Define the mapping of characters to their percent-encoded values for terraform
	encodingMap := map[rune]string{
		'%':  "%25",
		'/':  "%2F",
		'<':  "%3C",
		'>':  "%3E",
		'?':  "%3F",
		'[':  "%5B",
		'\\': "%5C",
		']':  "%5D",
		'^':  "%5E",
		'#':  "%23",
		'"':  "%22",
	}

	var sb strings.Builder
	// Loop through every charactor and replace with the encoded char
	for _, char := range input {
		if encoded, exists := encodingMap[char]; exists {
			sb.WriteString(encoded)
		} else {
			sb.WriteRune(char)
		}
	}
	return sb.String()
}

func main() {
	input := "Hello %<world># testing / special chars?"
	encoded := EncodeSpecialCharsForTerraform(input)
	fmt.Println(encoded)
}
