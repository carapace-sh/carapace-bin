package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var catFileCmd = &cobra.Command{
	Use:     "cat-file",
	Short:   "Provide content or type and size information for repository objects",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(catFileCmd).Standalone()

	catFileCmd.Flags().Bool("allow-unknown-type", false, "allow -s and -t to work with broken/corrupt objects")
	catFileCmd.Flags().String("batch", "", "show info and content of objects fed from the standard input")
	catFileCmd.Flags().Bool("batch-all-objects", false, "show all objects with --batch or --batch-check")
	catFileCmd.Flags().String("batch-check", "", "show info about objects fed from the standard input")
	catFileCmd.Flags().Bool("buffer", false, "buffer --batch output")
	catFileCmd.Flags().BoolS("e", "e", false, "exit with zero when there's no error")
	catFileCmd.Flags().String("filter", "", "omit objects from the list of printed objects (batch mode)")
	catFileCmd.Flags().Bool("filters", false, "for blob objects, run filters on object's content")
	catFileCmd.Flags().Bool("follow-symlinks", false, "follow in-tree symlinks (used with --batch or --batch-check)")
	catFileCmd.Flags().Bool("mailmap", false, "use mailmap file to map author/committer names")
	catFileCmd.Flags().Bool("no-filter", false, "do not omit objects from the list of printed objects")
	catFileCmd.Flags().Bool("no-mailmap", false, "do not use mailmap file")
	catFileCmd.Flags().Bool("no-use-mailmap", false, "do not use mailmap file")
	catFileCmd.Flags().BoolS("p", "p", false, "pretty-print object's content")
	catFileCmd.Flags().String("path", "", "use a specific path for --textconv/--filters")
	catFileCmd.Flags().BoolS("s", "s", false, "show object size")
	catFileCmd.Flags().BoolS("t", "t", false, "show object type")
	catFileCmd.Flags().Bool("textconv", false, "for blob objects, run textconv on object's content")
	catFileCmd.Flags().Bool("unordered", false, "do not order --batch-all-objects output")
	catFileCmd.Flags().Bool("use-mailmap", false, "use mailmap file to map author/committer names")
	rootCmd.AddCommand(catFileCmd)

	carapace.Gen(catFileCmd).FlagCompletion(carapace.ActionMap{
		"filter": git.ActionObjectFilters(),
	})
	carapace.Gen(catFileCmd).PositionalCompletion(
		carapace.ActionValues("blob", "tree", "commit", "tag"),
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(catFileCmd).DashAnyCompletion(
		carapace.ActionPositional(catFileCmd),
	)
}
