
package mdparse

import (
	"strings"
)

const (
	boldAnsi      = "\033[1m"
	underlineAnsi = "\033[4m"
	resetAnsi     = "\033[0m"
	block         = "█"
)

func Parse(content string) string {
	var result strings.Builder
	length := len(content)

	boldActive := false
	underlineActive := false

	for i := 0; i < length; i++ {
		char := content[i]

				nextChar := byte(0)
		if i+1 < length {
			nextChar = content[i+1]
		}

		switch char {
		case '*', '_': 			if nextChar == char { 				if hasClosing(content, i, "**") {
					boldActive = !boldActive
					if boldActive {
						result.WriteString(boldAnsi)
					} else {
						result.WriteString(resetAnsi)
					}
					i++ 				} else {
					result.WriteByte(char)
				}
			} else { 				if hasClosing(content, i, string(char)) {
					underlineActive = !underlineActive
					if underlineActive {
						result.WriteString(underlineAnsi)
					} else {
						result.WriteString(resetAnsi)
					}
				} else {
					result.WriteByte(char)
				}
			}

		case '>': 			if nextChar == ' ' {
				result.WriteString(block + " ")
				i++ 			} else if nextChar == '>' { 				result.WriteString(block + block)
				i++
			} else {
				result.WriteByte(char)
			}

		default:
			result.WriteByte(char)
		}
	}

		result.WriteString(resetAnsi)

	return result.String()
}

func hasClosing(content string, start int, marker string) bool {
	return strings.Contains(content[start+len(marker):], marker)
}
