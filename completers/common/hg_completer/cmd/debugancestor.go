package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugancestorCmd = &cobra.Command{
	Use:    "debugancestor",
	Short:  "find the ancestor revision of two revisions in a given index",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugancestorCmd).Standalone()

	rootCmd.AddCommand(debugancestorCmd)
}
