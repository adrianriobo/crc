package cmd

import (
	"fmt"
	"time"

	clicumber "github.com/code-ready/clicumber/testsuite"
)

const (
	retryWait                 = "60s"
	retryCount                = 15
	CRCExecutableInstalled    = "installed"
	CRCExecutableNotInstalled = "notInstalled"
)

func WaitForClusterInState(state string) error {
	retryDuration, err := time.ParseDuration(retryWait)
	if err != nil {
		return err
	}

	for i := 0; i < retryCount; i++ {
		err := CheckCRCStatus(state)
		if err == nil {
			return nil
		}
		time.Sleep(retryDuration)
	}
	return fmt.Errorf("cluster did not start properly")
}

func CheckCRCStatus(state string) error {
	expression := `.*OpenShift: .*Running \(v\d+\.\d+\.\d+.*\).*`
	if state == "stopped" {
		expression = ".*OpenShift: .*Stopped.*"
	}

	err := clicumber.ExecuteCommand("crc status")
	if err != nil {
		return err
	}
	return clicumber.CommandReturnShouldMatch("stdout", expression)
}

func CheckCRCExecutableState(state string) error {
	command := "which crc"
	switch state {
	case CRCExecutableInstalled:
		return clicumber.ExecuteCommandSucceedsOrFails(command, "succeeds")
	case CRCExecutableNotInstalled:
		return clicumber.ExecuteCommandSucceedsOrFails(command, "fails")
	default:
		return fmt.Errorf("%s state is not defined as valid crc executable state", state)
	}
}

func CheckMachineNotExists() error {
	expression := `.*Machine does not exist.*`
	err := clicumber.ExecuteCommand("crc status")
	if err != nil {
		return err
	}
	return clicumber.CommandReturnShouldMatch("stderr", expression)
}
