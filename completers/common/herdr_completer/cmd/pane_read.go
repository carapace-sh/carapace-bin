package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var pane_readCmd = &cobra.Command{
	Use:   "read <pane_id>",
	Short: "Read pane terminal output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_readCmd).Standalone()

	pane_readCmd.Flags().Bool("ansi", false, "")
	pane_readCmd.Flags().String("format", "", "")
	pane_readCmd.Flags().String("lines", "", "")
	pane_readCmd.Flags().Bool("raw", false, "")
	pane_readCmd.Flags().String("source", "", "")
	paneCmd.AddCommand(pane_readCmd)

	carapace.Gen(pane_readCmd).PositionalCompletion(herdr.ActionPanes(herdr.PaneOpts{}))

	carapace.Gen(pane_readCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("text", "ansi"),
		"source": carapace.ActionValues("visible", "recent", "recent-unwrapped", "detection"),
	})
}
