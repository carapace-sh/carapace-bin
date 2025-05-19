package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var mr_diffCmd = &cobra.Command{
	Use:   "diff [<id> | <branch>]",
	Short: "View changes in a merge request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_diffCmd).Standalone()

	mr_diffCmd.Flags().String("color", "", "Use color in diff output: always, never, auto.")
	mr_diffCmd.Flags().Bool("raw", false, "Use raw diff format that can be piped to commands")
	mrCmd.AddCommand(mr_diffCmd)

	carapace.Gen(mr_diffCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
	})

	carapace.Gen(mr_diffCmd).PositionalAnyCompletion(
		action.ActionMergeRequestsAndBranches(mr_diffCmd, ""),
	)
}
