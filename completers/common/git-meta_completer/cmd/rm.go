package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git-meta_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a string metadata key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()

	rmCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).PositionalCompletion(
		action.ActionTarget(),
	)
}
