package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Send actions to a specific session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_actionCmd).Standalone()

	helpCmd.AddCommand(help_actionCmd)
}
