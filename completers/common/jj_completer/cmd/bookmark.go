package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "Manage bookmarks [default alias: b]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmarkCmd).Standalone()

	bookmarkCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(bookmarkCmd)
}
