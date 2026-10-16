package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustCmd = &cobra.Command{
	Use:   "trust",
	Short: "Marks a config file as trusted",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustCmd).Standalone()

	trustCmd.Flags().BoolP("all", "a", false, "Trust all config files in workspace")
	trustCmd.Flags().BoolP("untrust", "u", false, "Untrust specified file")
	rootCmd.AddCommand(trustCmd)

	carapace.Gen(trustCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
