package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log [OPTIONS] [PATHS]...",
	Short: "Show commit history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()

	logCmd.Flags().Bool("color-words", false, "Show a word-level diff with changes indicated only by color")
	logCmd.Flags().Bool("git", false, "Show a Git-format diff")
	logCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	logCmd.Flags().StringP("limit", "n", "", "Limit number of revisions to show")
	logCmd.Flags().Bool("no-graph", false, "Don't show the graph, show a flat list of revisions")
	logCmd.Flags().BoolP("patch", "p", false, "Show patch")
	logCmd.Flags().Bool("reversed", false, "Show revisions in the opposite order (older revisions first)")
	logCmd.Flags().StringSliceP("revisions", "r", []string{"@ | ancestors(immutable_heads().., 2) | heads(immutable_heads())"}, "Which revisions to show. Defaults to the `revsets.log` setting, or `@ | ancestors(immutable_heads().., 2) | heads(immutable_heads())` if it is not set")
	logCmd.Flags().Bool("stat", false, "Show a histogram of the changes")
	logCmd.Flags().BoolP("summary", "s", false, "For each path, show only whether it was modified, added, or removed")
	logCmd.Flags().StringP("template", "T", "", "Render each revision using the given template")
	logCmd.Flags().String("tool", "", "Generate diff by external command")
	logCmd.Flags().Bool("types", false, "For each path, show only its type before and after")
	rootCmd.AddCommand(logCmd)

	carapace.Gen(logCmd).FlagCompletion(carapace.ActionMap{
		"revisions": jj.ActionRevSets(jj.RevOption{}.Default()),
		"tool":      bridge.ActionCarapaceBin().Split(),
	})

	carapace.Gen(logCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			revisions, err := logCmd.Flags().GetStringSlice("revisions") // TODO use revsets.log if set and flag is not set
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return jj.ActionRevChanges(revisions...)
		}),
	)
}
