package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:    "prune",
	Short:  "Prune the serialized git tree, dropping old entries",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()

	pruneCmd.Flags().Bool("dry-run", false, "Show what would be pruned without committing")
	pruneCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(pruneCmd)
}
