package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rdepsCmd = &cobra.Command{
	Use:   "rdeps",
	Short: "Recursively list dependencies of a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdepsCmd).Standalone()
	rdepsCmd.Flags().Bool("full", false, "Show every dependency in full traversal")
	rdepsCmd.Flags().Bool("index", false, "Use cached port index")
	rdepsCmd.Flags().Bool("no-build", false, "Exclude build-time dependencies")
	rdepsCmd.Flags().Bool("no-test", false, "Exclude test dependencies")
	rootCmd.AddCommand(rdepsCmd)
}
