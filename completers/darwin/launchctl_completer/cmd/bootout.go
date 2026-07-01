package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootoutCmd = &cobra.Command{
	Use:   "bootout",
	Short: "Tears down a domain or removes a service from a domain",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(bootoutCmd).Standalone()
	rootCmd.AddCommand(bootoutCmd)
	carapace.Gen(bootoutCmd).PositionalAnyCompletion(
		actionDomainTargets(),
	)
}
