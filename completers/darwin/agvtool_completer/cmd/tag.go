package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create a tag with the current version number",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tagCmd).Standalone()
	tagCmd.Flags().BoolP("force", "F", false, "Force tagging even if there are uncommitted changes")
	tagCmd.Flags().BoolP("noupdatecheck", "Q", false, "Skip the update check")
	rootCmd.AddCommand(tagCmd)
}
