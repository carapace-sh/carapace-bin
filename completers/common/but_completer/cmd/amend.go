package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var amendCmd = &cobra.Command{
	Use:     "amend",
	Short:   "Amend uncommitted changes into a commit or branch",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "editing commits",
}

func init() {
	carapace.Gen(amendCmd).Standalone()

	amendCmd.Flags().Bool("allow-merged", false, "Allow targeting branches and commits that are already merged upstream")
	amendCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	amendCmd.Flags().StringP("target", "t", "", "The commit or branch to amend into")
	amendCmd.MarkFlagRequired("target")
	rootCmd.AddCommand(amendCmd)

	carapace.Gen(amendCmd).FlagCompletion(carapace.ActionMap{
		"target": carapace.Batch(
			but.ActionCommits(),
			but.ActionCliIds(but.CliIdsOpts{Commits: true}),
		).ToA(),
	})

	carapace.Gen(amendCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCliIds(but.CliIdsOpts{Changes: true}),
		).ToA(),
	)
}
