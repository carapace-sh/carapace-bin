package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push commits to remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().Bool("fast-forward-merge", false, "Allow the server to fast-forward merge if the target branch head has moved")
	pushCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).PositionalCompletion(
		action.ActionBranches(pushCmd),
	)
}
