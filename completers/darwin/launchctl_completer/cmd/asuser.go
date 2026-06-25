package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var asuserCmd = &cobra.Command{
	Use:   "asuser",
	Short: "Execute a program in the bootstrap context of a given user",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(asuserCmd).Standalone()
	rootCmd.AddCommand(asuserCmd)
	carapace.Gen(asuserCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
