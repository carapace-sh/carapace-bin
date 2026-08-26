package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var stash_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply the stashed changes to the working tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_applyCmd).Standalone()

	stash_applyCmd.Flags().Bool("index", false, "attempt to recreate the index")
	stash_applyCmd.Flags().String("label-base", "", "label for the base in diff3 conflict markers")
	stash_applyCmd.Flags().String("label-ours", "", "label for the upstream side in conflict markers")
	stash_applyCmd.Flags().String("label-theirs", "", "label for the stashed side in conflict markers")
	stash_applyCmd.Flags().BoolP("quiet", "q", false, "be quiet, only report errors")
	stashCmd.AddCommand(stash_applyCmd)

	carapace.Gen(stash_applyCmd).PositionalCompletion(
		git.ActionStashes(),
	)
}
