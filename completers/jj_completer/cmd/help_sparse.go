package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_sparseCmd = &cobra.Command{
	Use:   "sparse",
	Short: "Manage which paths from the working-copy commit are present in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_sparseCmd).Standalone()

	helpCmd.AddCommand(help_sparseCmd)
}
