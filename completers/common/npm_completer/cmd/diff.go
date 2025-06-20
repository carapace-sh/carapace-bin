package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "The registry diff command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()
	diffCmd.Flags().StringArray("diff", nil, "argument to compare")
	diffCmd.Flags().String("diff-dest-prefix", "", "destination prefix to be used in output")
	diffCmd.Flags().Bool("diff-ignore-all-space", false, "ignore whitespace when comparing lines")
	diffCmd.Flags().Bool("diff-name-only", false, "print only filenames")
	diffCmd.Flags().Bool("diff-no-prefix", false, "do not show any source or destination prefix")
	diffCmd.Flags().String("diff-src-prefix", "", "source prefix to be used in output")
	diffCmd.Flags().Bool("diff-text", false, "treat all files as text")
	diffCmd.Flags().Int("diff-unified", 3, "numver of lines of context to print")
	diffCmd.Flags().BoolP("global", "g", false, "operate globally")
	diffCmd.Flags().String("tag", "latest", "tag to use when version is omitted")

	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).FlagCompletion(carapace.ActionMap{
		"diff": carapace.Batch(
			carapace.ActionFiles(),
			action.ActionPackages(diffCmd).UnlessF(condition.CompletingPath),
		).ToA(),
	})
}
