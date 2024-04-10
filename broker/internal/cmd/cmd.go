package cmd

import (
	"errors"
	"fmt"
	"strings"
)

type Cmd struct {
	IdFrom  string
	IdTo    string
	Command string
	Content string
}

func New(idFrom, idTo, command, content string) *Cmd {
	return &Cmd{
		IdFrom:  idFrom,
		IdTo:    idTo,
		Command: command,
		Content: content,
	}
}

func Encode(data string) (*Cmd, error) {
	dataMap := strings.Split(strings.Replace(data, `\n\n`, "\n\n", -1), "\n\n")
	if len(dataMap) <= 1 {
		return nil, errors.New("o dado não possui header ou body")
	}

	header := strings.Split(strings.Replace(data, `\n`, "\n", -1), "\n")
	body := dataMap[1]

	if len(header) <= 2 ||
		!strings.Contains(header[0], "IdFrom: ") ||
		!strings.Contains(header[1], "IdTo: ") ||
		!strings.Contains(header[2], "Cmd: ") {
		return nil, errors.New("header inválido")
	}

	idFrom := strings.Split(header[0], "IdFrom: ")[1]
	idTo := strings.Split(header[1], "IdTo: ")[1]
	command := strings.Split(header[2], "Cmd: ")[1]
	content := body

	return &Cmd{
		IdFrom:  idFrom,
		IdTo:    idTo,
		Command: command,
		Content: content,
	}, nil
}

func (cmd *Cmd) Decode() string {
	return fmt.Sprintf(
		"IdFrom: %s\nIdTo: %s\nCmd: %s\n\n%s",
		cmd.IdFrom, cmd.IdTo, cmd.Command, cmd.Content,
	)
}
