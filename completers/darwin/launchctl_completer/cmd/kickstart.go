package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kickstartCmd = &cobra.Command{
	Use:   "kickstart",
	Short: "Forces an existing service to start",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(kickstartCmd).Standalone()
	rootCmd.AddCommand(kickstartCmd)
	carapace.Gen(kickstartCmd).PositionalAnyCompletion(
		actionDomainTargets(),
	)
}
