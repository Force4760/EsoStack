package lexer

// Split the src code into words
func (le *Lexer) Split() {
	buffer := ""

	isComment := 0
	for _, i := range le.input {
		if isComment != 0 {

			switch i {
			case '(':
				isComment += 1
			case ')':
				isComment -= 1
			}

		} else {

			switch i {
			case '(':
				isComment += 1
			case ' ', '\n', '\t':
				// Delimiters that DO NOT matter

				if buffer != "" {
					// append non buffers
					le.chrs = append(le.chrs, buffer)
					buffer = ""
				}

			case '{', '}':
				// Delimiters that DO matter

				if buffer != "" {
					// append non buffers
					le.chrs = append(le.chrs, buffer)
					buffer = ""
				}

				// append { or }
				le.chrs = append(le.chrs, string(i))

			default:
				// Any other rune
				buffer += string(i)
			}

		}

	}

	// Append the final word if it exists
	if buffer != "" {
		le.chrs = append(le.chrs, buffer)
	}
}
