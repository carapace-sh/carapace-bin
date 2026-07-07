package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_pickCmd = &cobra.Command{
	Use:   "pick",
	Short: "Cherry-pick a commit onto the current branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_pickCmd).Standalone()

	helpCmd.AddCommand(help_pickCmd)
}
