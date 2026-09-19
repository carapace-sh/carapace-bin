package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "show changed files in the working directory",
	Aliases: []string{"st"},
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().BoolP("added", "a", false, "show only added files")
	statusCmd.Flags().BoolP("all", "A", false, "show status of all files")
	statusCmd.Flags().String("change", "", "list the changed files of a revision")
	statusCmd.Flags().BoolP("clean", "c", false, "show only files without changes")
	statusCmd.Flags().BoolP("copies", "C", false, "show source of copied files (DEFAULT: ui.statuscopies)")
	statusCmd.Flags().BoolP("deleted", "d", false, "show only missing files")
	statusCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	statusCmd.Flags().BoolP("ignored", "i", false, "show only ignored files")
	statusCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	statusCmd.Flags().BoolP("modified", "m", false, "show only modified files")
	statusCmd.Flags().BoolP("no-status", "n", false, "hide status prefix")
	statusCmd.Flags().BoolP("print0", "0", false, "end filenames with NUL, for use with xargs")
	statusCmd.Flags().BoolP("removed", "r", false, "show only removed files")
	statusCmd.Flags().String("rev", "", "show difference from revision")
	statusCmd.Flags().BoolP("subrepos", "S", false, "recurse into subrepositories")
	statusCmd.Flags().StringP("template", "T", "", "display with template")
	statusCmd.Flags().StringP("terse", "t", "", "show the terse output (EXPERIMENTAL)")
	statusCmd.Flags().BoolP("unknown", "u", false, "show only unknown (not tracked) files")
	rootCmd.AddCommand(statusCmd)

	carapace.Gen(statusCmd).FlagCompletion(carapace.ActionMap{
		"change":  action.ActionRevisions(),
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
		"rev":     action.ActionRevisions(),
	})

	carapace.Gen(statusCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
