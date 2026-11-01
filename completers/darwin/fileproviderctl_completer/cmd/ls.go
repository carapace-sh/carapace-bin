package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "run an interactive enumeration of the specified provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsCmd).Standalone()
	rootCmd.AddCommand(lsCmd)
	carapace.Gen(lsCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
