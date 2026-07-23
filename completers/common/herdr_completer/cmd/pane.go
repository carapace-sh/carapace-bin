package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paneCmd = &cobra.Command{
	Use:   "pane",
	Short: "Control terminal panes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paneCmd).Standalone()

	rootCmd.AddCommand(paneCmd)
}
