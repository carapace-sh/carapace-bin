package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Commands for creating and publishing code reviews to a forge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_reviewCmd).Standalone()

	helpCmd.AddCommand(help_reviewCmd)
}
