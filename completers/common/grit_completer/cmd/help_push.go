package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Publish the current branch to its remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_pushCmd).Standalone()

	helpCmd.AddCommand(help_pushCmd)
}
