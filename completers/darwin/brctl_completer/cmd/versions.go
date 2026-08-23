package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionsCmd)
}

var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "list non-local versions of a document",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionsCmd).Standalone()

	versionsCmd.Flags().Bool("all", false, "List all non-local versions including those that are locally cached")
}