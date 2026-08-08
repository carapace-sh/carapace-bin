package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Check if your `.gitlab-ci.yml` file is valid.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_lintCmd).Standalone()

	ci_lintCmd.Flags().Bool("dry-run", false, "Run pipeline creation simulation.")
	ci_lintCmd.Flags().Bool("include-jobs", false, "Include the list of jobs that would exist in a static check or pipeline simulation.")
	ci_lintCmd.Flags().String("ref", "", "When '--dry-run' is true, sets the branch or tag context for validating the CI/CD YAML configuration.")
	ciCmd.AddCommand(ci_lintCmd)

	carapace.Gen(ci_lintCmd).FlagCompletion(carapace.ActionMap{
		"ref": carapace.Batch(
			action.ActionBranches(ci_lintCmd),
			action.ActionTags(ci_lintCmd),
		).ToA(),
	})

	carapace.Gen(ci_lintCmd).PositionalCompletion(
		carapace.ActionFiles(".gitlab-ci.yml", ".gitlab-ci.yaml", ".yml", ".yaml"),
	)
}
