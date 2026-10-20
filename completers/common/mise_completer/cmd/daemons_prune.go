// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "[experimental] Remove daemon state left by deleted projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_pruneCmd).Standalone()

	daemons_pruneCmd.Flags().BoolP("dry-run", "n", false, "Show what would be removed without deleting anything")
	daemons_pruneCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonsCmd.AddCommand(daemons_pruneCmd)
}
