package installer

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	clicumber "github.com/code-ready/clicumber/testsuite"
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
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		return string(path.Dir(filename) +
			string(filepath.Separator) +
			scriptsRelativePath), nil
	}
	return "", fmt.Errorf("error recovering required resources for applescript installer handler")
}

func (a applescriptHandler) Install() error {
	command := strings.Join(append([]string{"osascript"},
		installScript, *a.InstallerPath,
		*a.CurrentUserPassword), " ")
	return clicumber.ExecuteCommandSucceedsOrFails(command, "succeeds")
}
