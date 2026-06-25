package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unloadCmd = &cobra.Command{
	Use:   "unload",
	Short: "Unloads a service or directory of services",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(unloadCmd).Standalone()
	rootCmd.AddCommand(unloadCmd)
	carapace.Gen(unloadCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
