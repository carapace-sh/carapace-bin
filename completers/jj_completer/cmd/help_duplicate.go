package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_duplicateCmd = &cobra.Command{
	Use:   "duplicate",
	Short: "Create a new change with the same content as an existing one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_duplicateCmd).Standalone()

	helpCmd.AddCommand(help_duplicateCmd)
}
