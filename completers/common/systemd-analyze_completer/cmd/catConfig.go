package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var catConfigCmd = &cobra.Command{
	Use:   "cat-config",
	Short: "Show configuration file and drop-ins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catConfigCmd).Standalone()

	rootCmd.AddCommand(catConfigCmd)

	carapace.Gen(catConfigCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
