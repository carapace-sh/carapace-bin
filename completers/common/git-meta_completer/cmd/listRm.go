package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git-meta_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listRmCmd = &cobra.Command{
	Use:   "list:rm",
	Short: "Show list entries with IDs, or remove one by index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listRmCmd).Standalone()

	listRmCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(listRmCmd)

	carapace.Gen(listRmCmd).PositionalCompletion(
		action.ActionTarget(),
	)
}
