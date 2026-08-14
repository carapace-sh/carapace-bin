package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_setPaneBorderlessCmd = &cobra.Command{
	Use:   "set-pane-borderless",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_setPaneBorderlessCmd).Standalone()

	action_helpCmd.AddCommand(action_help_setPaneBorderlessCmd)
}
