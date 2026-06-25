package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the specified service if it is running",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(stopCmd).Standalone()
	rootCmd.AddCommand(stopCmd)
	carapace.Gen(stopCmd).PositionalAnyCompletion(
		actionDomainTargets(),
	)
}
