package execcommand_test

import (
	"testing"

	execcommand "github.com/mvstermind/file-watcher/execCommand"
)

func TestRun(t *testing.T) {
	t.Run("string without flag", func(t *testing.T) {

		emptyString := []string{"string", "with", "no", "cmd"}
		value := execcommand.Run(emptyString)

		if value != "" {
			t.Errorf("string without flag found value of %v, when nothing was expected", value)
		}

	})

	t.Run("string with flag", func(t *testing.T) {

		validArg := []string{"-e", "echo", "hello", "world"}

		value := execcommand.Run(validArg)

		if value != "echo hello world" {
			t.Errorf("expected value was `echo hello world` but got %v", value)

		}
	})

}
