package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:     "export",
	Short:   "dump the header and diffs for one or more changesets",
	GroupID: groups[group_change_import_export].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exportCmd).Standalone()

	exportCmd.Flags().Bool("binary", false, "generate binary diffs in git mode (default)")
	exportCmd.Flags().StringP("bookmark", "B", "", "export changes only reachable by given bookmark")
	exportCmd.Flags().BoolP("git", "g", false, "use git extended diff format (DEFAULT: diff.git)")
	exportCmd.Flags().Bool("nodates", false, "omit dates from diff headers")
	exportCmd.Flags().StringP("output", "o", "", "print output to file with formatted name")
	exportCmd.Flags().StringArrayP("rev", "r", nil, "revisions to export")
	exportCmd.Flags().Bool("switch-parent", false, "diff against the second parent")
	exportCmd.Flags().StringP("template", "T", "", "display with template")
	exportCmd.Flags().BoolP("text", "a", false, "treat all files as text")
	rootCmd.AddCommand(exportCmd)

	carapace.Gen(exportCmd).FlagCompletion(carapace.ActionMap{
		"bookmark": hg.ActionBookmarks(),
		"output":   carapace.ActionFiles(),
		"rev":      hg.ActionRevisions(),
	})

	carapace.Gen(exportCmd).PositionalAnyCompletion(
		hg.ActionRevisions(),
	)
}
