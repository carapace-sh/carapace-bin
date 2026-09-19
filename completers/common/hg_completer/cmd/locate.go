package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var locateCmd = &cobra.Command{
	Use:     "locate",
	Short:   "locate files matching specific patterns (DEPRECATED)",
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(locateCmd).Standalone()

	locateCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	locateCmd.Flags().BoolP("fullpath", "f", false, "print complete paths from the filesystem root")
	locateCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	locateCmd.Flags().BoolP("print0", "0", false, "end filenames with NUL, for use with xargs")
	locateCmd.Flags().StringP("rev", "r", "", "search the repository as it is in REV")
	rootCmd.AddCommand(locateCmd)

	carapace.Gen(locateCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
		"rev":     action.ActionRevisions(),
	})

	carapace.Gen(locateCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
