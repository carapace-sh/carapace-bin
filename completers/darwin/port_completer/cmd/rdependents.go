package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rdependentsCmd = &cobra.Command{
	Use:   "rdependents",
	Short: "Recursively list installed ports that depend on a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdependentsCmd).Standalone()
	rdependentsCmd.Flags().Bool("full", false, "Show every dependent in full traversal")
	rootCmd.AddCommand(rdependentsCmd)
}
