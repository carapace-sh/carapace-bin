package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var commitTreeCmd = &cobra.Command{
	Use:     "commit-tree",
	Short:   "Create a new commit object",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(commitTreeCmd).Standalone()

	commitTreeCmd.Flags().StringS("F", "F", "", "read commit log message from file")
	commitTreeCmd.Flags().StringP("gpg-sign", "S", "", "GPG sign commit")
	commitTreeCmd.Flags().StringS("m", "m", "", "commit message")
	commitTreeCmd.Flags().Bool("no-gpg-sign", false, "do not GPG sign commit")
	commitTreeCmd.Flags().StringSliceS("p", "p", nil, "id of a parent commit object")
	rootCmd.AddCommand(commitTreeCmd)

	commitTreeCmd.Flag("gpg-sign").NoOptDefVal = " "

	carapace.Gen(commitTreeCmd).FlagCompletion(carapace.ActionMap{
		"F":        carapace.ActionFiles(),
		"gpg-sign": os.ActionGpgKeyIds(),
		"p":        git.ActionRefs(git.RefOption{}.Default()), // TODO what are valid parents?
	})

	carapace.Gen(commitTreeCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()), // TODO what's a tree object?
	)
}
