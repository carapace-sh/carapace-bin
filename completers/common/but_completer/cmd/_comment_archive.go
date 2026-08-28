package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var _comment_archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive a comment, hiding it from all future listings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(_comment_archiveCmd).Standalone()

	_comment_archiveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	_commentCmd.AddCommand(_comment_archiveCmd)
}