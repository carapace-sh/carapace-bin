package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var metaeditCmd = &cobra.Command{
	Use:   "metaedit [OPTIONS] [REVSETS]...",
	Short: "Modify the metadata of a revision without changing its content",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(metaeditCmd).Standalone()

	metaeditCmd.Flags().String("author", "", "Set author to the provided string")
	metaeditCmd.Flags().String("author-timestamp", "", "Set the author date to the given date either human readable, eg Sun, 23 Jan 2000 01:23:45 JST) or as a time stamp, eg 2000-01-23T01:23:45+09:00)")
	metaeditCmd.Flags().Bool("force-rewrite", false, "Rewrite the commit, even if no other metadata changed")
	metaeditCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	metaeditCmd.Flags().StringSliceP("message", "m", nil, "Update the change description")
	metaeditCmd.Flags().Bool("update-author", false, "Update the author to the configured user")
	metaeditCmd.Flags().Bool("update-author-timestamp", false, "Update the author timestamp")
	metaeditCmd.Flags().Bool("update-change-id", false, "Generate a new change-id")
	rootCmd.AddCommand(metaeditCmd)

	carapace.Gen(metaeditCmd).PositionalAnyCompletion(
		jj.ActionRevSets(jj.RevOption{}.Default()),
	)
}
