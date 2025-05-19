package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var verifyCommitCmd = &cobra.Command{
	Use:     "verify-commit",
	Short:   "Check the GPG signature of commits",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(verifyCommitCmd).Standalone()

	verifyCommitCmd.Flags().Bool("raw", false, "print raw gpg status output")
	verifyCommitCmd.Flags().BoolP("verbose", "v", false, "print commit contents")
	rootCmd.AddCommand(verifyCommitCmd)

	carapace.Gen(verifyCommitCmd).PositionalAnyCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
