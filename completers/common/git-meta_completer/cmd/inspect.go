package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git-meta_completer/cmd/action"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Browse metadata keys and values",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectCmd).Standalone()

	inspectCmd.Flags().BoolP("help", "h", false, "Print help")
	inspectCmd.Flags().Bool("promisor", false, "List only promisor (not-yet-fetched) keys")
	inspectCmd.Flags().Bool("timeline", false, "Show a weekly timeline graph of entries")
	rootCmd.AddCommand(inspectCmd)

	carapace.Gen(inspectCmd).PositionalCompletion(
		action.ActionTarget(),
		carapace.ActionValues(),
	)
}
