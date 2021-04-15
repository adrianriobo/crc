// +build !windows

package installer

import (
	"runtime"
	"strings"

	clicumber "github.com/code-ready/clicumber/testsuite"
	applescript "github.com/code-ready/crc/test/extended/os/applescript"
)

const (
	scriptsRelativePath string = "scripts"
	installScript       string = "install.applescript"
)

type applescriptHandler struct {
	CurrentUserPassword *string
	InstallerPath       *string
}

func NewInstaller(currentUserPassword *string, installerPath *string) Installer {
	// TODO check parameters as they are mandatory otherwise exit
	if runtime.GOOS == "darwin" {
		return applescriptHandler{
			CurrentUserPassword: currentUserPassword,
			InstallerPath:       installerPath}

	}
	return nil
}

func RequiredResourcesPath() (string, error) {
	return applescript.GetScriptsPath(scriptsRelativePath)
}

func (a applescriptHandler) Install() error {
	command := strings.Join(append([]string{"osascript"},
		installScript, *a.InstallerPath,
		*a.CurrentUserPassword), " ")
	return clicumber.ExecuteCommandSucceedsOrFails(command, "succeeds")
}
