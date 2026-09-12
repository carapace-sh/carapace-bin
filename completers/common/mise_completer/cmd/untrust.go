package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var untrustCmd = &cobra.Command{
	Use:   "untrust",
	Short: "Marks a config file as untrusted",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(untrustCmd).Standalone()

	rootCmd.AddCommand(untrustCmd)

	carapace.Gen(untrustCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
