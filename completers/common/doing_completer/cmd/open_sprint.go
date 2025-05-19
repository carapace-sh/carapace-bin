package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var open_sprintCmd = &cobra.Command{
	Use:   "sprint",
	Short: "Open current sprint view",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(open_sprintCmd).Standalone()

	open_sprintCmd.Flags().Bool("help", false, "Show this message and exit.")
	openCmd.AddCommand(open_sprintCmd)
}
