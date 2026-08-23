package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var enumerateCmd = &cobra.Command{
	Use:   "enumerate",
	Short: "run an interactive enumeration of the specified provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enumerateCmd).Standalone()
	rootCmd.AddCommand(enumerateCmd)
	carapace.Gen(enumerateCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
