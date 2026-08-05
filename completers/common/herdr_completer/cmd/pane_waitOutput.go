package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var pane_waitOutputCmd = &cobra.Command{
	Use:   "wait-output <pane_id>",
	Short: "Wait for matching pane output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_waitOutputCmd).Standalone()

	pane_waitOutputCmd.Flags().String("lines", "", "Restrict the searched snapshot to N lines")
	pane_waitOutputCmd.Flags().String("match", "", "Match a literal substring")
	pane_waitOutputCmd.Flags().Bool("raw", false, "Keep ANSI escape sequences while matching")
	pane_waitOutputCmd.Flags().String("regex", "", "Match a Rust regular expression")
	pane_waitOutputCmd.Flags().String("source", "", "Terminal snapshot source (default: recent)")
	pane_waitOutputCmd.Flags().String("timeout", "", "Fail after this many milliseconds")
	paneCmd.AddCommand(pane_waitOutputCmd)

	carapace.Gen(pane_waitOutputCmd).PositionalCompletion(herdr.ActionPanes(herdr.PaneOpts{}))

	carapace.Gen(pane_waitOutputCmd).FlagCompletion(carapace.ActionMap{
		"source": carapace.ActionValues("visible", "recent", "recent-unwrapped"),
	})
}
