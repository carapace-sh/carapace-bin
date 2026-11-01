package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git-meta_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listPopCmd = &cobra.Command{
	Use:   "list:pop",
	Short: "Pop a value from a list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listPopCmd).Standalone()

	listPopCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(listPopCmd)

	carapace.Gen(listPopCmd).PositionalCompletion(
		action.ActionTarget(),
	)
}
