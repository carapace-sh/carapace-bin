package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stash_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save local modifications to a new stash (deprecated, use push)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_saveCmd).Standalone()

	stash_saveCmd.Flags().BoolP("keep-index", "k", false, "keep changed added to index")
	stash_saveCmd.Flags().StringP("message", "m", "", "set description")
	stash_saveCmd.Flags().BoolP("patch", "p", false, "interactively select hunks between HEAD and working tree")
	stash_saveCmd.Flags().BoolP("quiet", "q", false, "suppress feedback messages")
	stash_saveCmd.Flags().BoolP("staged", "S", false, "only stash staged")
	stashCmd.AddCommand(stash_saveCmd)
}
