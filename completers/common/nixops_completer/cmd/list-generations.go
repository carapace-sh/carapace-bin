package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ListGenerationsCmd = &cobra.Command{
	Use:   "list-generations",
	Short: "List Generations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ListGenerationsCmd).Standalone()
	rootCmd.AddCommand(ListGenerationsCmd)
}
