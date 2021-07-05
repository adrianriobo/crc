// +build !windows

package cmd

import "fmt"

func (c Command) ToString() string {
	cmd := ""
	if c.DisableUpdateCheck {
		cmd += "CRC_DISABLE_UPDATE_CHECK=true "
	}
	if c.DisableNTP {
		cmd += "CRC_DEBUG_ENABLE_STOP_NTP=true "
	}
	return cmd + fmt.Sprintf("crc  %s", c.Action)
}
