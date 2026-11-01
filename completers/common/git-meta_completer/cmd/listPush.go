package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git-meta_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listPushCmd = &cobra.Command{
	Use:   "list:push",
	Short: "Push a value onto a list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listPushCmd).Standalone()

	listPushCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(listPushCmd)

	carapace.Gen(listPushCmd).PositionalCompletion(
		action.ActionTarget(),
	)
}
