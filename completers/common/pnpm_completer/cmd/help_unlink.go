package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_unlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Removes links to a local package and reinstalls it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_unlinkCmd).Standalone()

	helpCmd.AddCommand(help_unlinkCmd)
}
