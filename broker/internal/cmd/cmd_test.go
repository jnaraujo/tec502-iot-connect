package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeCmd(t *testing.T) {
	cmd := Cmd{
		IdFrom:  "BROKER",
		IdTo:    "test_id",
		Command: "test_command",
		Content: "test_content",
	}

	expected := "IdFrom: BROKER\nIdTo: test_id\nCmd: test_command\n\ntest_content"
	got := cmd.Decode()

	assert.Equal(t, got, expected)
}

func TestEncode(t *testing.T) {
	data := "IdFrom: test_id_from\nIdTo: test_id_to\nCmd: test_command\n\ntest_content"

	expected := Cmd{
		IdFrom:  "test_id_from",
		IdTo:    "test_id_to",
		Command: "test_command",
		Content: "test_content",
	}

	got, err := Encode(data)

	assert.NoError(t, err)

	assert.Equal(t, got.Command, expected.Command)
	assert.Equal(t, got.Content, expected.Content)
}
