package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Update the description and create a new change on top",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_commitCmd).Standalone()

	helpCmd.AddCommand(help_commitCmd)
}
