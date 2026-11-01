package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var blameCmd = &cobra.Command{
	Use:     "blame",
	Short:   "Show who last modified each line of a file",
	GroupID: "history",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blameCmd).Standalone()

	blameCmd.Flags().Bool("first-parent", false, "At merge commits, follow only the first parent (Git's `--first-parent`)")
	blameCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	blameCmd.Flags().StringSlice("ignore-rev", nil, "Ignore this revision when attributing lines (repeatable)")
	blameCmd.Flags().StringSlice("ignore-revs-file", nil, "Read revision oids to ignore from a file (repeatable; `\"\"` clears prior entries)")
	blameCmd.Flags().StringP("lines", "L", "", "Limit to an inclusive 1-based line range (`START,END`)")
	blameCmd.Flags().String("rev", "", "Blame at this revision instead of `HEAD`")
	rootCmd.AddCommand(blameCmd)

	carapace.Gen(blameCmd).FlagCompletion(carapace.ActionMap{
		"rev": git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(blameCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
