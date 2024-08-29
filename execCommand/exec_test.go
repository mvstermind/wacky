package execcommand_test

import (
	"testing"

	execcommand "github.com/mvstermind/wacky/execCommand"
)

func TestRun(t *testing.T) {
	t.Run("No flag provided", func(t *testing.T) {
		args := []string{"string", "with", "no", "cmd"}
		value := execcommand.Run(args)

		if value != "" {
			t.Errorf("Expected an empty string, but got '%v'", value)
		}
	})

	t.Run("Flag with valid command", func(t *testing.T) {
		args := []string{"-e", "echo", "hello", "world"}
		value := execcommand.Run(args)

		expected := "echo hello world"
		if value != expected {
			t.Errorf("Expected '%v', but got '%v'", expected, value)
		}
	})

	t.Run("Flag without command", func(t *testing.T) {
		args := []string{"-e"}
		value := execcommand.Run(args)

		if value != "" {
			t.Errorf("Expected an empty string, but got '%v'", value)
		}
	})

	t.Run("Flag at the end without command", func(t *testing.T) {
		args := []string{"some", "other", "args", "-e"}
		value := execcommand.Run(args)

		if value != "" {
			t.Errorf("Expected an empty string, but got '%v'", value)
		}
	})

	t.Run("Multiple flags, only the first is processed", func(t *testing.T) {
		args := []string{"some", "other", "args", "-e", "echo", "hello", "-e", "echo", "world"}
		value := execcommand.Run(args)

		expected := "echo hello -e echo world"
		if value != expected {
			t.Errorf("Expected '%v', but got '%v'", expected, value)
		}
	})

	t.Run("Flag with complex command", func(t *testing.T) {
		args := []string{"-e", "ls", "-la", "/home/user"}
		value := execcommand.Run(args)

		expected := "ls -la /home/user"
		if value != expected {
			t.Errorf("Expected '%v', but got '%v'", expected, value)
		}
	})

	t.Run("No arguments provided", func(t *testing.T) {
		args := []string{}
		value := execcommand.Run(args)

		if value != "" {
			t.Errorf("Expected an empty string, but got '%v'", value)
		}
	})

	t.Run("Only command flag provided", func(t *testing.T) {
		args := []string{"-e"}
		value := execcommand.Run(args)

		if value != "" {
			t.Errorf("Expected an empty string, but got '%v'", value)
		}
	})

	t.Run("Long flag --execute", func(t *testing.T) {
		args := []string{"--execute", "echo", "hello", "world"}
		value := execcommand.Run(args)

		expected := "echo hello world"
		if value != expected {
			t.Errorf("Expected '%v', but got '%v'", expected, value)
		}
	})

	t.Run("Long flag with additional arguments", func(t *testing.T) {
		args := []string{"--execute", "echo", "hello", "--some-flag"}
		value := execcommand.Run(args)

		expected := "echo hello --some-flag"
		if value != expected {
			t.Errorf("Expected '%v', but got '%v'", expected, value)
		}
	})
}
