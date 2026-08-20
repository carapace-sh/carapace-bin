package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_dumpLayoutCmd = &cobra.Command{
	Use:   "dump-layout",
	Short: "Dump current layout to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_dumpLayoutCmd).Standalone()

	action_dumpLayoutCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_dumpLayoutCmd)
}
