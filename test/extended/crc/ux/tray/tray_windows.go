// +build windows

package tray

import (
	"fmt"
)

type TrayAutoitHandler struct {
	BundleLocation *string
	PullSecretFile *string
}

func NewTray(bundleLocationValue *string, pullSecretFileValue *string) Tray {
	return TrayAutoitHandler{
		BundleLocation: bundleLocationValue,
		PullSecretFile: pullSecretFileValue}

}

func RequiredResourcesPath() (string, error) {
	return "", nil
}

func (t TrayApplescriptHandler) Install() error {
	return fmt.Errorf("not implemented yet")
}

func (t TrayApplescriptHandler) IsInstalled() error {
	return fmt.Errorf("not implemented yet")
}

func (t TrayApplescriptHandler) IsAccessible() error {
	return fmt.Errorf("not implemented yet")
}

func (t TrayApplescriptHandler) ClickStart() error {
	return fmt.Errorf("not implemented yet")
}

func (t TrayApplescriptHandler) ClickStop() error {
	return fmt.Errorf("not implemented yet")
}

func (t TrayApplescriptHandler) ClickDelete() error {
	return fmt.Errorf("not implemented yet")
}

func (t TrayApplescriptHandler) ClickQuit() error {
	return fmt.Errorf("not implemented yet")
}

func (t TrayApplescriptHandler) SetPullSecretFile() error {
	return fmt.Errorf("not implemented yet")
}

func (t TrayApplescriptHandler) IsClusterRunning() error {
	return fmt.Errorf("not implemented yet")
}

func (t TrayApplescriptHandler) IsClusterStopped() error {
	return fmt.Errorf("not implemented yet")
}

func (t TrayApplescriptHandler) CopyOCLoginCommandAsKubeadmin() error {
	return fmt.Errorf("not implemented yet")
}

func (t TrayApplescriptHandler) CopyOCLoginCommandAsDeveloper() error {
	return fmt.Errorf("not implemented yet")
}

func (t TrayApplescriptHandler) ConnectClusterAsKubeadmin() error {
	return fmt.Errorf("not implemented yet")
}

func (t TrayApplescriptHandler) ConnectClusterAsDeveloper() error {
	return fmt.Errorf("not implemented yet")
}
