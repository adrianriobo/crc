// +build windows

package tray

import (
	"fmt"
)

type autoitHandler struct {
	BundleLocation *string
	PullSecretFile *string
}

func NewTray(bundleLocationValue *string, pullSecretFileValue *string) Tray {
	return autoitHandler{
		BundleLocation: bundleLocationValue,
		PullSecretFile: pullSecretFileValue}

}

func RequiredResourcesPath() (string, error) {
	return "", nil
}

func (a autoitHandler) Install() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) IsInstalled() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) IsAccessible() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) ClickStart() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) ClickStop() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) ClickDelete() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) ClickQuit() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) SetPullSecretFile() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) IsClusterRunning() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) IsClusterStopped() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) CopyOCLoginCommandAsKubeadmin() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) CopyOCLoginCommandAsDeveloper() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) ConnectClusterAsKubeadmin() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) ConnectClusterAsDeveloper() error {
	return fmt.Errorf("not implemented yet")
}
