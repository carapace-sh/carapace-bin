package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List running browsers and their tabs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsCmd).Standalone()

	lsCmd.Flags().Bool("all", false, "Every browser, not just this terminal tab")
	lsCmd.Flags().Bool("json", false, "Machine readable, including cdp ports and pane ids")
	rootCmd.AddCommand(lsCmd)
}
