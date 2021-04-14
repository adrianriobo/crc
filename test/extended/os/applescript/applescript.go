package applescript

import (
	"strings"

	clicumber "github.com/code-ready/clicumber/testsuite"
)

func ExecuteApplescript(scriptFilename string, args ...string) error {
	command := strings.Join(append(
		append([]string{"osascript"}, scriptFilename),
		args...),
		" ")
	return clicumber.ExecuteCommandSucceedsOrFails(command, "succeeds")
}

func ExecuteApplescriptReturnShouldMatch(expectedOutput string,
	scriptFilename string, args ...string) error {
	command := strings.Join(append(
		append([]string{"osascript"}, scriptFilename),
		args...),
		" ")
	err := clicumber.ExecuteCommand(command)
	if err != nil {
		return err
	}
	return clicumber.CommandReturnShouldMatch("stdout", expectedOutput)
}
