package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var mergeFileCmd = &cobra.Command{
	Use:     "merge-file",
	Short:   "Run a three-way file merge",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(mergeFileCmd).Standalone()

	mergeFileCmd.Flags().StringSliceS("L", "L", nil, "use labels in place of the file names in conflict reports")
	mergeFileCmd.Flags().String("diff-algorithm", "", "use a different diff algorithm while merging")
	mergeFileCmd.Flags().Bool("diff3", false, "show conflicts in \"diff3\" style")
	mergeFileCmd.Flags().Bool("object-id", false, "specify the contents to merge as blobs in the current repository instead of files")
	mergeFileCmd.Flags().Bool("ours", false, "resolve conflicts favouring our side of the lines")
	mergeFileCmd.Flags().BoolS("p", "p", false, "send results to standard output instead of overwriting <current>")
	mergeFileCmd.Flags().BoolS("q", "q", false, "do not warn about conflicts")
	mergeFileCmd.Flags().Bool("theirs", false, "resolve conflicts favouring their side of the lines")
	mergeFileCmd.Flags().Bool("union", false, "resolve conflicts favouring both side of the lines")
	mergeFileCmd.Flags().Bool("zdiff3", false, "show conflicts in \"zdiff3\" style")
	rootCmd.AddCommand(mergeFileCmd)

	carapace.Gen(mergeFileCmd).FlagCompletion(carapace.ActionMap{
		"diff-algorithm": git.ActionDiffAlgorithms(),
	})

	carapace.Gen(mergeFileCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
