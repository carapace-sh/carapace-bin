package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alias_rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Short:   "Remove an alias using its hostname",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_rmCmd).Standalone()

	alias_rmCmd.Flags().Bool("yes", false, "Skip confirmation")

	aliasCmd.AddCommand(alias_rmCmd)

	carapace.Gen(alias_rmCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
