package installer

import (
	"runtime"

	"github.com/code-ready/crc/test/extended/crc/ux/installer/darwin/applescript"
)

type Installer interface {
	Install() error
}

func NewInstaller(currentUserPassword *string, installerPath *string) Installer {
	// TODO check parameters as they are mandatory otherwise exit
	if runtime.GOOS == "darwin" {
		return applescript.InstallerApplescriptHandler{
			CurrentUserPassword: currentUserPassword,
			InstallerPath:       installerPath}

	}
	return nil
}

func RequiredResourcesPath() (string, error) {
	if runtime.GOOS == "darwin" {
		return applescript.RequiredResourcesPath()
	}
	return "", nil
}
