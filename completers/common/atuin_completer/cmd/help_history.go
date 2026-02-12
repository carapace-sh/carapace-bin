package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Manipulate shell history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_historyCmd).Standalone()

	helpCmd.AddCommand(help_historyCmd)
}
