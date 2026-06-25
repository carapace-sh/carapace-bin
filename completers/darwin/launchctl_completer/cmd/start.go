package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the specified service",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()
	rootCmd.AddCommand(startCmd)
	carapace.Gen(startCmd).PositionalAnyCompletion(
		actionDomainTargets(),
	)
}
