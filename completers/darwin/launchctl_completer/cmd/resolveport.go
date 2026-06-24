package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resolveportCmd = &cobra.Command{
	Use:   "resolveport",
	Short: "Resolves a port name from a process to an endpoint in launchd",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(resolveportCmd).Standalone()
	rootCmd.AddCommand(resolveportCmd)
}
