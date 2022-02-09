package colorize

import "strings"

// Colorize a full string
func Colorize(s string, c Color) string {
	// "ColorCode CurrentString EndCode"
	return string(c) + s + string(END)
}

// Colorize a word
func ColorizeWord(s string, word string, c Color) string {
	// "ColorCode word EndCode"
	replacer := string(c) + word + string(END)

	replaced := strings.ReplaceAll(s, word, replacer)

	return replaced
}
