package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Configures the next invocation of a service for debugging",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(debugCmd).Standalone()
	rootCmd.AddCommand(debugCmd)
	carapace.Gen(debugCmd).PositionalAnyCompletion(
		actionDomainTargets(),
	)
}
