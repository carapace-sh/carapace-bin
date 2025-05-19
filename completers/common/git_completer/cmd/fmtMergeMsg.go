package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var fmtMergeMsgCmd = &cobra.Command{
	Use:     "fmt-merge-msg",
	Short:   "Produce a merge commit message",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(fmtMergeMsgCmd).Standalone()

	fmtMergeMsgCmd.Flags().StringP("file", "F", "", "Take the list of merged objects from <file> instead of stdin")
	fmtMergeMsgCmd.Flags().String("into-name", "", "Prepare the merge message as if merging to the branch <branch>")
	fmtMergeMsgCmd.Flags().String("log", "", "At most <n> commits from each merge parent will be used (20 if <n> is omitted)")
	fmtMergeMsgCmd.Flags().StringP("message", "m", "", "Use <message> instead of the branch names for the first line of the log message")
	fmtMergeMsgCmd.Flags().Bool("no-log", false, "Do not list one-line descriptions from the actual commits being merged")
	rootCmd.AddCommand(fmtMergeMsgCmd)

	carapace.Gen(fmtMergeMsgCmd).FlagCompletion(carapace.ActionMap{
		"file":      carapace.ActionFiles(),
		"into-name": git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true, Tags: true}),
	})
}
