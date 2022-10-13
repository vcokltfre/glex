package glex

import (
	"fmt"
	"strings"
)

func SplitCommand(command string) ([]string, error) {
	rawParts := []string{}

	current := ""
	escaped := false
	quoted := ' '

	for i, c := range command {
		char := string(c)

		if escaped {
			escaped = false
			switch c {
			case 'n':
				current += "\n"
			case 'r':
				current += "\r"
			case 't':
				current += "\t"
			case 'b':
				current += "\b"
			case 'f':
				current += "\f"
			case 'v':
				current += "\v"
			case '\\':
				current += "\\"
			case '"':
				current += "\""
			case '\'':
				current += "'"
			default:
				return nil, fmt.Errorf("invalid escape sequence \\%s at position %d", char, i)
			}
			continue
		}

		if c == '\\' {
			escaped = true
			continue
		}

		if quoted != ' ' {
			if c == quoted {
				rawParts = append(rawParts, current)
				quoted = ' '
				current = ""
				continue
			}

			current += char
			continue
		}

		if c == '"' || c == '\'' {
			quoted = c
			continue
		}

		if c == ' ' {
			if current != "" {
				rawParts = append(rawParts, current)
				current = ""
			}
			continue
		}

		current += char
	}

	if quoted != ' ' {
		return nil, fmt.Errorf("unterminated quote at position %d", len(command))
	}

	if escaped {
		return nil, fmt.Errorf("unterminated escape at position %d", len(command))
	}

	if current != "" {
		rawParts = append(rawParts, current)
	}

	parts := []string{}

	for _, part := range rawParts {
		if strings.HasPrefix(part, "--") && strings.Contains(part, "=") {
			parts = append(parts, part[:strings.Index(part, "=")])
			parts = append(parts, part[strings.Index(part, "=")+1:])
			continue
		}

		parts = append(parts, part)
	}

	return parts, nil
}
