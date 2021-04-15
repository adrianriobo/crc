// +build windows

package notification

import (
	"fmt"
)

type NotificationAutoitHandler struct {
}

func NewNotification() Notification {
	return NotificationAutoitHandler{}
}

func RequiredResourcesPath() (string, error) {
	return "", nil
}

func (n NotificationAutoitHandler) GetClusterRunning() error {
	return fmt.Errorf("not implemented yet")
}

func (n NotificationAutoitHandler) GetClusterStopped() error {
	return fmt.Errorf("not implemented yet")

}

func (n NotificationAutoitHandler) GetClusterDeleted() error {
	return fmt.Errorf("not implemented yet")
}

func (n NotificationAutoitHandler) ClearNotifications() error {
	return fmt.Errorf("not implemented yet")
}
