package cmd

import (
	"testing"
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

	if got != expected {
		t.Error("Error encoding cmd")
	}
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

	if err != nil {
		t.Error("Error decoding cmd")
	}

	if got.Command != expected.Command || got.Content != expected.Content {
		t.Error("Error decoding cmd - wrong data")
	}
}
