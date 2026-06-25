package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var blameCmd = &cobra.Command{
	Use:   "blame",
	Short: "Prints the reason a service is running",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(blameCmd).Standalone()
	rootCmd.AddCommand(blameCmd)
	carapace.Gen(blameCmd).PositionalAnyCompletion(
		actionDomainTargets(),
	)
}
