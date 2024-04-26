package utils

import (
	"strings"
)

func CenterFormat(msg string, width int) string {
	spaces := (width - len(msg)) / 2
	return strings.Repeat(" ", spaces) + msg
}
