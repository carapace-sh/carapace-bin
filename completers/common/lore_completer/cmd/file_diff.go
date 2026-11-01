package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var file_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show differences between two revisions of a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_diffCmd).Standalone()

	file_diffCmd.Flags().StringP("context", "U", "3", "Number of unchanged context lines to show around each hunk")
	file_diffCmd.Flags().Bool("diff3", false, "If given, produce three-way merge output with conflict markers instead of a two-way unified diff")
	file_diffCmd.Flags().BoolP("help", "h", false, "Print help")
	file_diffCmd.Flags().Bool("ignore-space-at-eol", false, "Treat lines that differ only in trailing whitespace as unchanged")
	file_diffCmd.Flags().Bool("ignore-space-change", false, "Collapse runs of internal whitespace to a single space before comparing")
	file_diffCmd.Flags().String("source", "", "Optional signature of the source revision to diff from, by default the current revision")
	file_diffCmd.Flags().String("target", "", "Optional signature of the target revision to diff to, by default the current file system state")
	file_diffCmd.Flags().String("targets", "", "Path to a targets file containing all the paths to all files")
	fileCmd.AddCommand(file_diffCmd)

	carapace.Gen(file_diffCmd).FlagCompletion(carapace.ActionMap{
		"source":  action.ActionRevisions(file_diffCmd),
		"target":  action.ActionRevisions(file_diffCmd),
		"targets": carapace.ActionFiles(),
	})

	carapace.Gen(file_diffCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
