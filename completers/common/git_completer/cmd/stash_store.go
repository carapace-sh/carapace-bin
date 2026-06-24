package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var stash_storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store a given stash entry in the reflog",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_storeCmd).Standalone()

	stash_storeCmd.Flags().StringP("message", "m", "", "set description")
	stash_storeCmd.Flags().BoolP("quiet", "q", false, "suppress feedback messages")
	stashCmd.AddCommand(stash_storeCmd)

	carapace.Gen(stash_storeCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
