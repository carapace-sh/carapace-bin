package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_squashCmd = &cobra.Command{
	Use:   "squash",
	Short: "Move changes from a revision into its parent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_squashCmd).Standalone()

	helpCmd.AddCommand(help_squashCmd)
}
