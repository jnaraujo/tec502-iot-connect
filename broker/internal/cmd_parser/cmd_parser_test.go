package cmd_parser

import (
	"testing"
)

func TestEncodeCmd(t *testing.T) {
	cmd := Cmd{
		ID:      "test_id",
		Command: "test_command",
		Content: "test_content",
	}

	expected := "Id: test_id\nCmd: test_command\n\ntest_content"
	got := EncodeCmd(cmd)

	if got != expected {
		t.Error("Error encoding cmd")
	}
}

func TestDecodeCmd(t *testing.T) {
	data := "Id: test_id\nCmd: test_command\n\ntest_content"

	expected := Cmd{
		ID:      "test_id",
		Command: "test_command",
		Content: "test_content",
	}

	got, err := DecodeCmd(data)

	if err != nil {
		t.Error("Error decoding cmd")
	}

	if got.Command != expected.Command || got.Content != expected.Content {
		t.Error("Error decoding cmd - wrong data")
	}
}
