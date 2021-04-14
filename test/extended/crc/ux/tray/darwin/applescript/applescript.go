package applescript

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	clicumber "github.com/code-ready/clicumber/testsuite"
	applescriptHelper "github.com/code-ready/crc/test/extended/os/applescript"
)

type Element struct {
	Name         string
	AXIdentifier string
}

type TrayApplescriptHandler struct {
	BundleLocation *string
	PullSecretFile *string
}

func RequiredResourcesPath() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		return filepath.Join(path.Dir(filename),
			scriptsRelativePath), nil
	}
	return "", fmt.Errorf("error recovering required resources for applescript tray handler")
}

func (t TrayApplescriptHandler) Install() error {
	err := clicumber.ExecuteCommandSucceedsOrFails("crc setup", "succeeds")
	if err != nil {
		return err
	}
	// Required to pass parameters with spaces to applescript
	sanitizedAppPath := strings.Join(append([]string{"\""}, appPath, "\""), "")
	return applescriptHelper.ExecuteApplescript(installTray, sanitizedAppPath)
}

func (t TrayApplescriptHandler) IsInstalled() error {
	return executeCommandSucceeds("launchctl list | grep crc", "0.*tray")
}

func (t TrayApplescriptHandler) IsAccessible() error {
	return checkAccessible(func() error {
		return applescriptHelper.ExecuteApplescript(
			checkTrayIconIsVisible, bundleIdentifier)
	}, "Tray icon")
}

func (t TrayApplescriptHandler) ClickStart() error {
	return clickButtonByAction(actionStart)
}

func (t TrayApplescriptHandler) ClickStop() error {
	return clickButtonByAction(actionStop)
}

func (t TrayApplescriptHandler) ClickDelete() error {
	return clickButtonByAction(actionDelete)
}

func (t TrayApplescriptHandler) ClickQuit() error {
	return clickButtonByAction(actionQuit)
}

func (t TrayApplescriptHandler) SetPullSecretFile() error {
	return applescriptHelper.ExecuteApplescript(
		setPullSecretFile, bundleIdentifier, *t.PullSecretFile)
}

func (t TrayApplescriptHandler) IsClusterRunning() error {
	return checkTrayShowsFieldWithValue(fieldState, stateRunning)
}

func (t TrayApplescriptHandler) IsClusterStopped() error {
	return checkTrayShowsFieldWithValue(fieldState, stateStopped)
}

func (t TrayApplescriptHandler) CopyOCLoginCommandAsKubeadmin() error {
	return clickCopyOCLoginCommand(userKubeadmin)
}

func (t TrayApplescriptHandler) CopyOCLoginCommandAsDeveloper() error {
	return clickCopyOCLoginCommand(userDeveloper)
}

func (t TrayApplescriptHandler) ConnectClusterAsKubeadmin() error {
	return applescriptHelper.ExecuteApplescriptReturnShouldMatch(
		userKubeadmin, runOCLoginCommand)
}

func (t TrayApplescriptHandler) ConnectClusterAsDeveloper() error {
	return applescriptHelper.ExecuteApplescriptReturnShouldMatch(
		userDeveloper, runOCLoginCommand)
}

func clickButtonByAction(actionName string) error {
	return clickOnElement(actionName, clickTrayMenuItem)
}

func clickCopyOCLoginCommand(userName string) error {
	return clickOnElement(userName, getOCLoginCommand)
}

func clickOnElement(elementName string, scriptName string) error {
	element, err := getElement(elementName, elements)
	if err != nil {
		return err
	}
	return applescriptHelper.ExecuteApplescript(
		scriptName, bundleIdentifier, element.AXIdentifier)
}

func checkTrayShowsFieldWithValue(field string, expectedValue string) error {
	element, err := getElement(field, elements)
	if err != nil {
		return err
	}
	return applescriptHelper.ExecuteApplescriptReturnShouldMatch(
		expectedValue, getTrayFieldlValue, bundleIdentifier, element.AXIdentifier)
}

func getElement(name string, elements []Element) (Element, error) {
	for _, e := range elements {
		if name == e.Name {
			return e, nil
		}
	}
	return Element{},
		fmt.Errorf("element '%s', Can not be accessed from the tray", name)
}

func checkAccessible(uxIsAccessible func() error, component string) error {
	retryDuration, err := time.ParseDuration(uxCheckAccessibilityDuration)
	if err != nil {
		return err
	}
	for i := 0; i < uxCheckAccessibilityRetry; i++ {
		err := uxIsAccessible()
		if err == nil {
			return nil
		}
		time.Sleep(retryDuration)
	}
	return fmt.Errorf("%s is not accessible", component)
}

// TODO review which helper use
func executeCommandSucceeds(command string, expectedOutput string) error {
	err := clicumber.ExecuteCommand(command)
	if err != nil {
		return err
	}
	return clicumber.CommandReturnShouldMatch("stdout", expectedOutput)
}
