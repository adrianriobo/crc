package notification

import (
	"runtime"

	"github.com/code-ready/crc/test/extended/crc/ux/notification/darwin/applescript"
)

type Notification interface {
	GetClusterRunning() error
	GetClusterStopped() error
	GetClusterDeleted() error
	ClearNotifications() error
}

func NewNotification() Notification {
	if runtime.GOOS == "darwin" {
		return applescript.NotificationApplescriptHandler{}

	}
	return nil
}

func RequiredResourcesPath() (string, error) {
	if runtime.GOOS == "darwin" {
		return applescript.RequiredResourcesPath()
	}
	return "", nil
}
