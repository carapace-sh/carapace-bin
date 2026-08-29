package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit changes from Git back to the p4 repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submitCmd).Standalone()

	submitCmd.Flags().BoolP("M", "M", false, "Detect renames")
	submitCmd.Flags().String("branch", "", "After submitting, sync this named branch instead of the default p4/master")
	submitCmd.Flags().String("commit", "", "Submit only the specified commit or range of commits")
	submitCmd.Flags().String("conflict", "", "Conflict behavior when applying a commit to p4")
	submitCmd.Flags().Bool("disable-p4sync", false, "Disable the automatic sync of p4/master after submit")
	submitCmd.Flags().Bool("disable-rebase", false, "Disable the automatic rebase after submit")
	submitCmd.Flags().BoolP("dry-run", "n", false, "Show what commits would be submitted to p4")
	submitCmd.Flags().Bool("export-labels", false, "Export tags from Git as p4 labels")
	submitCmd.Flags().Bool("no-verify", false, "Bypass p4-pre-submit and p4-changelist hooks")
	submitCmd.Flags().String("origin", "", "Upstream location from which commits are identified to submit")
	submitCmd.Flags().Bool("prepare-p4-only", false, "Apply a commit to the p4 workspace without submitting")
	submitCmd.Flags().Bool("preserve-user", false, "Re-author p4 changes before submitting to p4")
	submitCmd.Flags().Bool("shelve", false, "Create a series of shelved changelists instead of submitting")
	submitCmd.Flags().IntSlice("update-shelve", nil, "Update an existing shelved changelist with this commit")
	rootCmd.AddCommand(submitCmd)

	carapace.Gen(submitCmd).FlagCompletion(carapace.ActionMap{
		"branch":   carapace.ActionValues(),
		"commit":   carapace.ActionValues(),
		"conflict": carapace.ActionValues("ask", "skip", "quit"),
		"origin":   carapace.ActionValues(),
	})

	carapace.Gen(submitCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
