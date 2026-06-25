package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Bootstraps a service or directory of services",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(loadCmd).Standalone()
	rootCmd.AddCommand(loadCmd)
	carapace.Gen(loadCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
