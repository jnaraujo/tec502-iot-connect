package utils

import (
	"strings"
)

// CenterFormat centraliza uma string em um espaço de tamanho width.
func CenterFormat(msg string, width int) string {
	spaces := (width - len(msg)) / 2
	return strings.Repeat(" ", spaces) + msg
}
