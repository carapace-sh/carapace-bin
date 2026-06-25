package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstraps a domain or a service into a domain",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(bootstrapCmd).Standalone()
	rootCmd.AddCommand(bootstrapCmd)
	carapace.Gen(bootstrapCmd).PositionalCompletion(
		actionDomainTargets(),
	)
	carapace.Gen(bootstrapCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".plist"),
	)
}
