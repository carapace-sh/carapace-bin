package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pane_edgesCmd = &cobra.Command{
	Use:   "edges",
	Short: "Show pane edge information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_edgesCmd).Standalone()

	pane_edgesCmd.Flags().Bool("current", false, "")
	pane_edgesCmd.Flags().String("pane", "", "")
	paneCmd.AddCommand(pane_edgesCmd)
}
