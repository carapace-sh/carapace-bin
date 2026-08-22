package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_setPaneColorCmd = &cobra.Command{
	Use:   "set-pane-color",
	Short: "Set the default foreground/background color of a pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_setPaneColorCmd).Standalone()

	help_actionCmd.AddCommand(help_action_setPaneColorCmd)
}
