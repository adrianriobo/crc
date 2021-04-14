package tray

import (
	"runtime"

	"github.com/code-ready/crc/test/extended/crc/ux/tray/darwin/applescript"
)

type Tray interface {
	Install() error
	IsInstalled() error
	IsAccessible() error
	ClickStart() error
	ClickStop() error
	ClickDelete() error
	ClickQuit() error
	SetPullSecretFile() error
	IsClusterRunning() error
	IsClusterStopped() error
	CopyOCLoginCommandAsKubeadmin() error
	CopyOCLoginCommandAsDeveloper() error
	// TODO check if make sense create a new ux component
	ConnectClusterAsKubeadmin() error
	ConnectClusterAsDeveloper() error
}

func NewTray(bundleLocationValue *string, pullSecretFileValue *string) Tray {
	if runtime.GOOS == "darwin" {
		return applescript.TrayApplescriptHandler{
			BundleLocation: bundleLocationValue,
			PullSecretFile: pullSecretFileValue}

	}
	return nil
}

func RequiredResourcesPath() (string, error) {
	if runtime.GOOS == "darwin" {
		return applescript.RequiredResourcesPath()
	}
	return "", nil
}
