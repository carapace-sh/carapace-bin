package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:     "purge",
	Short:   "removes files not tracked by Mercurial",
	Aliases: []string{"clean"},
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(purgeCmd).Standalone()

	purgeCmd.Flags().BoolP("abort-on-err", "a", false, "abort if an error occurs")
	purgeCmd.Flags().Bool("all", false, "purge ignored files too")
	purgeCmd.Flags().Bool("confirm", false, "ask before permanently deleting files")
	purgeCmd.Flags().Bool("dirs", false, "purge empty directories")
	purgeCmd.Flags().StringArrayP("exclude", "X", nil, "exclude names matching the given patterns")
	purgeCmd.Flags().Bool("files", false, "purge files")
	purgeCmd.Flags().BoolP("ignored", "i", false, "purge only ignored files")
	purgeCmd.Flags().StringArrayP("include", "I", nil, "include names matching the given patterns")
	purgeCmd.Flags().BoolP("print", "p", false, "print filenames instead of deleting them")
	purgeCmd.Flags().BoolP("print0", "0", false, "end filenames with NUL, for use with xargs (implies -p/--print)")
	rootCmd.AddCommand(purgeCmd)

	carapace.Gen(purgeCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionFiles(),
		"include": carapace.ActionFiles(),
	})

	carapace.Gen(purgeCmd).PositionalAnyCompletion(
		action.ActionUntrackedFiles(),
	)
}
