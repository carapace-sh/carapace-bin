package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_dumpLayoutCmd = &cobra.Command{
	Use:   "dump-layout",
	Short: "Dump current layout to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_dumpLayoutCmd).Standalone()

	help_actionCmd.AddCommand(help_action_dumpLayoutCmd)
}
