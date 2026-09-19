package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:     "cat",
	Short:   "output the current or given revision of files",
	GroupID: groups[group_file_content_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()

	catCmd.Flags().Bool("decode", false, "apply any matching decode filter")
	catCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	catCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	catCmd.Flags().StringP("output", "o", "", "print output to file with formatted name")
	catCmd.Flags().StringP("rev", "r", "", "print the given revision")
	catCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(catCmd)

	carapace.Gen(catCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
		"output":  carapace.ActionFiles(),
		"rev":     action.ActionRevisions(),
	})

	carapace.Gen(catCmd).PositionalAnyCompletion(
		action.ActionTrackedFiles(),
	)
}
