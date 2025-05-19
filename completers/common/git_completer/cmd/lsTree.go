package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var lsTreeCmd = &cobra.Command{
	Use:     "ls-tree",
	Short:   "List the contents of a tree object",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(lsTreeCmd).Standalone()

	lsTreeCmd.Flags().String("abbrev", "", "use <n> digits to display object names")
	lsTreeCmd.Flags().BoolS("d", "d", false, "only show trees")
	lsTreeCmd.Flags().String("format", "", "format to use for the output")
	lsTreeCmd.Flags().Bool("full-name", false, "use full path names")
	lsTreeCmd.Flags().Bool("full-tree", false, "list entire tree")
	lsTreeCmd.Flags().BoolP("long", "l", false, "include object size")
	lsTreeCmd.Flags().Bool("name-only", false, "list only filenames")
	lsTreeCmd.Flags().Bool("name-status", false, "list only filenames")
	lsTreeCmd.Flags().Bool("no-abbrev", false, "do not use abbreviated digits to display object names")
	lsTreeCmd.Flags().Bool("no-full-name", false, "do not use full path names")
	lsTreeCmd.Flags().Bool("no-full-tree", false, "do not list entire tree")
	lsTreeCmd.Flags().Bool("object-only", false, "list only objects")
	lsTreeCmd.Flags().BoolS("r", "r", false, "recurse into subtrees")
	lsTreeCmd.Flags().BoolS("t", "t", false, "show trees when recursing")
	lsTreeCmd.Flags().BoolS("z", "z", false, "terminate entries with NUL byte")
	rootCmd.AddCommand(lsTreeCmd)

	lsTreeCmd.Flag("abbrev").NoOptDefVal = " "

	carapace.Gen(lsTreeCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(lsTreeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRefFiles(c.Args[0])
		}),
	)
}
