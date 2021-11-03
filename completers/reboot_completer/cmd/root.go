package cmd

import (
	"github.com/rsteube/carapace-bin/completers/halt_completer/cmd"
)

/**
Description for go:generate
	Use: "reboot",
	Short: "reboot the machine",
*/

func Execute() error {
	return cmd.ExecuteReboot()
}
