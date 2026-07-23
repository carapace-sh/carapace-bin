package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var pane_splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split a pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_splitCmd).Standalone()

	pane_splitCmd.Flags().Bool("current", false, "")
	pane_splitCmd.Flags().String("cwd", "", "")
	pane_splitCmd.Flags().String("direction", "", "")
	pane_splitCmd.Flags().StringSlice("env", nil, "Set an environment variable for the launched process")
	pane_splitCmd.Flags().Bool("focus", false, "")
	pane_splitCmd.Flags().Bool("no-focus", false, "")
	pane_splitCmd.Flags().String("pane", "", "")
	pane_splitCmd.Flags().String("ratio", "", "")
	paneCmd.AddCommand(pane_splitCmd)

	carapace.Gen(pane_splitCmd).FlagCompletion(carapace.ActionMap{
		"cwd":       carapace.ActionFiles(),
		"direction": carapace.ActionValues("right", "down"),
		"pane":      herdr.ActionPanes(herdr.PaneOpts{}),
	})

	carapace.Gen(pane_splitCmd).PositionalCompletion(herdr.ActionPanes(herdr.PaneOpts{}))
}
