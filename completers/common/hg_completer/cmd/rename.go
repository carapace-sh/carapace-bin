package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:     "rename",
	Short:   "rename files; equivalent of copy + remove",
	Aliases: []string{"move", "mv"},
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(renameCmd).Standalone()

	renameCmd.Flags().BoolP("after", "A", false, "record a rename that has already occurred")
	renameCmd.Flags().String("at-rev", "", "(un)mark renames in the given revision (EXPERIMENTAL)")
	renameCmd.Flags().BoolP("dry-run", "n", false, "do not perform actions, just print output")
	renameCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	renameCmd.Flags().BoolP("force", "f", false, "forcibly move over an existing managed file")
	renameCmd.Flags().Bool("forget", false, "unmark a destination file as renamed")
	renameCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	rootCmd.AddCommand(renameCmd)

	carapace.Gen(renameCmd).FlagCompletion(carapace.ActionMap{
		"at-rev":  hg.ActionRevisions(),
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
	})

	carapace.Gen(renameCmd).PositionalCompletion(
		hg.ActionTrackedFiles(),
		carapace.ActionFiles(),
	)
}
