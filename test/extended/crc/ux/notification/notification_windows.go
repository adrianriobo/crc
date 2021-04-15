// +build windows

package notification

import (
	"fmt"
)

type autoitHandler struct {
}

func NewNotification() Notification {
	return autoitHandler{}
}

func RequiredResourcesPath() (string, error) {
	return "", nil
}

func (a autoitHandler) GetClusterRunning() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) GetClusterStopped() error {
	return fmt.Errorf("not implemented yet")

}

func (a autoitHandler) GetClusterDeleted() error {
	return fmt.Errorf("not implemented yet")
}

func (a autoitHandler) ClearNotifications() error {
	return fmt.Errorf("not implemented yet")
}
