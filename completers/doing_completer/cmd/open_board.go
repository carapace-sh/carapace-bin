package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var open_boardCmd = &cobra.Command{
	Use:   "board",
	Short: "Open board view",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(open_boardCmd).Standalone()

	open_boardCmd.Flags().Bool("help", false, "Show this message and exit.")
	openCmd.AddCommand(open_boardCmd)
}
