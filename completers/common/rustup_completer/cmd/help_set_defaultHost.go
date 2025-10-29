package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_set_defaultHostCmd = &cobra.Command{
	Use:   "default-host",
	Short: "The triple used to identify toolchains when not specified",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_set_defaultHostCmd).Standalone()

	help_setCmd.AddCommand(help_set_defaultHostCmd)
}
