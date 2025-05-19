package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var open_pipeCmd = &cobra.Command{
	Use:   "pipe",
	Short: "Open latest pipeline runs for repository view",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(open_pipeCmd).Standalone()

	open_pipeCmd.Flags().Bool("help", false, "Show this message and exit.")
	openCmd.AddCommand(open_pipeCmd)
}
