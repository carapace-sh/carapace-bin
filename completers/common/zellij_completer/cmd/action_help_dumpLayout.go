package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_dumpLayoutCmd = &cobra.Command{
	Use:   "dump-layout",
	Short: "Dump current layout to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_dumpLayoutCmd).Standalone()

	action_helpCmd.AddCommand(action_help_dumpLayoutCmd)
}
