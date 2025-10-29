package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var set_help_defaultHostCmd = &cobra.Command{
	Use:   "default-host",
	Short: "The triple used to identify toolchains when not specified",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_help_defaultHostCmd).Standalone()

	set_helpCmd.AddCommand(set_help_defaultHostCmd)
}
