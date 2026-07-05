package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "List dependencies of a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(depsCmd).Standalone()
	depsCmd.Flags().Bool("index", false, "Use cached port index instead of Portfile")
	depsCmd.Flags().Bool("no-build", false, "Exclude build-time dependencies")
	depsCmd.Flags().Bool("no-test", false, "Exclude test dependencies")
	rootCmd.AddCommand(depsCmd)
}
