package cmd

import (
	"testing"
)

func TestDecodeCmd(t *testing.T) {
	cmd := Cmd{
		ID:      "test_id",
		Command: "test_command",
		Content: "test_content",
	}

	expected := "Id: test_id\nCmd: test_command\n\ntest_content"
	got := cmd.Decode()

	if got != expected {
		t.Error("Error encoding cmd")
	}
}

func TestEncode(t *testing.T) {
	data := "Id: test_id\nCmd: test_command\n\ntest_content"

	expected := Cmd{
		ID:      "test_id",
		Command: "test_command",
		Content: "test_content",
	}

	got, err := Encode(data)

	if err != nil {
		t.Error("Error decoding cmd")
	}

	if got.Command != expected.Command || got.Content != expected.Content {
		t.Error("Error decoding cmd - wrong data")
	}
}
