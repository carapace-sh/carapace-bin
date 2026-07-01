package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists information about services",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()
	rootCmd.AddCommand(listCmd)
	carapace.Gen(listCmd).PositionalAnyCompletion(
		actionDomainTargets(),
	)
}
