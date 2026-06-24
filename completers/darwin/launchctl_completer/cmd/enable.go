package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enables an existing service",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(enableCmd).Standalone()
	rootCmd.AddCommand(enableCmd)
	carapace.Gen(enableCmd).PositionalAnyCompletion(
		actionDomainTargets(),
	)
}
