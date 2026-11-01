package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show differences between two revisions of a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().StringP("context", "U", "3", "Number of unchanged context lines to show around each hunk")
	diffCmd.Flags().Bool("diff3", false, "If given, produce three-way merge output with conflict markers instead of a two-way unified diff")
	diffCmd.Flags().BoolP("help", "h", false, "Print help")
	diffCmd.Flags().Bool("ignore-space-at-eol", false, "Treat lines that differ only in trailing whitespace as unchanged")
	diffCmd.Flags().Bool("ignore-space-change", false, "Collapse runs of internal whitespace to a single space before comparing")
	diffCmd.Flags().String("source", "", "Optional signature of the source revision to diff from, by default the current revision")
	diffCmd.Flags().String("target", "", "Optional signature of the target revision to diff to, by default the current file system state")
	diffCmd.Flags().String("targets", "", "Path to a targets file containing all the paths to all files")
	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).FlagCompletion(carapace.ActionMap{
		"source":  action.ActionRevisions(diffCmd),
		"target":  action.ActionRevisions(diffCmd),
		"targets": carapace.ActionFiles(),
	})

	carapace.Gen(diffCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
