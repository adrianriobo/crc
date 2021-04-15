// +build !windows

package notification

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"time"

	applescriptHelper "github.com/code-ready/crc/test/extended/os/applescript"
)

type NotificationApplescriptHandler struct {
}

const (
	scriptsRelativePath           string = "scripts"
	manageNotifications           string = "manageNotifications.applescript"
	manageNotificationActionGet   string = "get"
	manageNotificationActionClear string = "clear"

	notificationDelay   string = "5s"
	notificationRetries int    = 36
)

func NewNotification() Notification {
	if runtime.GOOS == "darwin" {
		return NotificationApplescriptHandler{}

	}
	return nil
}

func RequiredResourcesPath() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		return filepath.Join(path.Dir(filename),
			scriptsRelativePath), nil
	}
	return "", fmt.Errorf("error recovering required resources for applescript notification handler")
}

func (n NotificationApplescriptHandler) GetClusterRunning() error {
	return checkNotificationMessage(startMessage)
}

func (n NotificationApplescriptHandler) GetClusterStopped() error {
	return checkNotificationMessage(stopMessage)

}

func (n NotificationApplescriptHandler) GetClusterDeleted() error {
	return checkNotificationMessage(deleteMessage)
}

func (n NotificationApplescriptHandler) ClearNotifications() error {
	return applescriptHelper.ExecuteApplescript(manageNotifications, manageNotificationActionClear)
}

func checkNotificationMessage(notificationMessage string) error {
	retryDelay, err := time.ParseDuration(notificationDelay)
	if err != nil {
		return err
	}
	for i := 0; i < notificationRetries; i++ {
		err := applescriptHelper.ExecuteApplescriptReturnShouldMatch(
			notificationMessage, manageNotifications, manageNotificationActionGet)
		if err == nil {
			return nil
		}
		time.Sleep(retryDelay)
	}
	return fmt.Errorf("notification: %s. Timeout", notificationMessage)
}
