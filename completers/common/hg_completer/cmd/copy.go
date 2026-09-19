package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:     "copy",
	Short:   "mark files as copied for the next commit",
	Aliases: []string{"cp"},
	GroupID: groups[group_file_content_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copyCmd).Standalone()

	copyCmd.Flags().BoolP("after", "A", false, "record a copy that has already occurred")
	copyCmd.Flags().String("at-rev", "", "(un)mark copies in the given revision (EXPERIMENTAL)")
	copyCmd.Flags().BoolP("dry-run", "n", false, "do not perform actions, just print output")
	copyCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	copyCmd.Flags().BoolP("force", "f", false, "forcibly copy over an existing managed file")
	copyCmd.Flags().Bool("forget", false, "unmark a destination file as copied")
	copyCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	rootCmd.AddCommand(copyCmd)

	carapace.Gen(copyCmd).FlagCompletion(carapace.ActionMap{
		"at-rev":  action.ActionRevisions(),
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
	})

	carapace.Gen(copyCmd).PositionalCompletion(
		action.ActionTrackedFiles(),
		carapace.ActionFiles(),
	)
}
