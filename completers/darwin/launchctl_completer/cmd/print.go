package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Prints a description of a domain or service",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(printCmd).Standalone()
	rootCmd.AddCommand(printCmd)
	carapace.Gen(printCmd).PositionalAnyCompletion(
		actionDomainTargets(),
	)
}
