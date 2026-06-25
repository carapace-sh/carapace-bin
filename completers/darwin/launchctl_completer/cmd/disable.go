package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disables an existing service",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(disableCmd).Standalone()
	rootCmd.AddCommand(disableCmd)
	carapace.Gen(disableCmd).PositionalAnyCompletion(
		actionDomainTargets(),
	)
}
