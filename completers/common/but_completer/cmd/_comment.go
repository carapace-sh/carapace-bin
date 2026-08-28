package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var _commentCmd = &cobra.Command{
	Use:    "_comment",
	Short:  "Work with ephemeral comments anchored to lines in diffs",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(_commentCmd).Standalone()

	_commentCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(_commentCmd)
}
