package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "list non-local versions of a document",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionsCmd).Standalone()
	rootCmd.AddCommand(versionsCmd)

	versionsCmd.Flags().Bool("all", false, "List all non-local versions including those that are locally cached")
}