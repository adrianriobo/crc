// +build !windows

package installer

import (
	"fmt"

	goautoit "github.com/shadow1163/goautoit"
)

type InstallerAutoitHandler struct {
	CurrentUserPassword *string
	InstallerPath       *string
}

func NewInstaller(currentUserPassword *string, installerPath *string) Installer {
	// TODO check parameters as they are mandatory otherwise exit
	return InstallerAutoitHandler{
		CurrentUserPassword: currentUserPassword,
		InstallerPath:       installerPath}
}

func RequiredResourcesPath() (string, error) {
	return "", nil
}

func (i InstallerAutoitHandler) Install() error {
	command := fmt.Sprintf("msiexec.exe /i %s /qf", *i.InstallerPath)
	installerPid := goautoit.RunWait(command)
	if installerPid == 0 {
		return fmt.Errorf("error starting the msi installer")
	}
	return nil
}
