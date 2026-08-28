package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var _comment_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all comments, with their anchors refreshed against the current diffs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(_comment_listCmd).Standalone()

	_comment_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	_comment_listCmd.Flags().String("timeout", "60", "How many seconds --wait blocks before giving up for this invocation")
	_comment_listCmd.Flags().Bool("wait", false, "Block until a comment exists instead of returning an empty listing")
	_commentCmd.AddCommand(_comment_listCmd)
}