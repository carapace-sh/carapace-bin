package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/git_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var stash_showCmd = &cobra.Command{
	Use:   "show",
	Short: "show the changes recorded in the stash entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_showCmd).Standalone()
	stash_showCmd.Flags().Bool("include-untracked", false, "include untracked")
	stash_showCmd.Flags().Bool("only-untracked", false, "only untracked")
	common.AddDiffFlags(stash_showCmd)

	stashCmd.AddCommand(stash_showCmd)

	carapace.Gen(stash_showCmd).PositionalCompletion(git.ActionStashes())
}
