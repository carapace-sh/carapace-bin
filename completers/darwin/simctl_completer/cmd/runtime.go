package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runtimeCmd = &cobra.Command{
	Use:   "runtime",
	Short: "Manage simulator runtime disk images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runtimeCmd).Standalone()
	rootCmd.AddCommand(runtimeCmd)
	carapace.Gen(runtimeCmd).PositionalCompletion(
		carapace.ActionValues("add", "delete"),
	)
	carapace.Gen(runtimeCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
