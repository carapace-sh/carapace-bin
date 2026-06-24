package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Unloads the specified service name",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()
	rootCmd.AddCommand(removeCmd)
	carapace.Gen(removeCmd).PositionalAnyCompletion(
		actionDomainTargets(),
	)
}
