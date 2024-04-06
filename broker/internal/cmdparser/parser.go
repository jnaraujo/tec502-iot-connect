package cmdparser

import (
	"broker/internal/errors"
	"fmt"
	"strings"
)

type Cmd struct {
	ID      string
	Command string
	Content string
}

func DecodeCmd(data string) (*Cmd, error) {
	dataMap := strings.Split(data, "\n\n")

	if len(dataMap) <= 1 {
		return nil, errors.ErrInvalidData
	}

	header := strings.Split(dataMap[0], "\n")

	if len(header) <= 1 || !strings.Contains(header[0], "Id: ") || !strings.Contains(header[1], "Cmd: ") {
		return nil, errors.ErrInvalidData
	}

	id := strings.Split(header[0], "Id: ")[1]
	content := dataMap[1]
	command := strings.Split(dataMap[0], "Cmd: ")[1]

	return &Cmd{
		ID:      id,
		Command: command,
		Content: content,
	}, nil
}

func EncodeCmd(cmd Cmd) string {
	return fmt.Sprintf("Id: %s\nCmd: %s\n\n%s", cmd.ID, cmd.Command, cmd.Content)
}
