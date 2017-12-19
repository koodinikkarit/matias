package ewdatabases

import (
	"strings"
)

func PrepareWords(original string) string {
	return `{\rtf1{\pard ` + strings.Replace(original, `\n`, `\par`, -1)
}
