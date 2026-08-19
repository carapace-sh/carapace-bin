package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listGenerationsCmd = &cobra.Command{
	Use:   "list-generations",
	Short: "list previous configurations for rollback",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listGenerationsCmd).Standalone()
	rootCmd.AddCommand(listGenerationsCmd)
}
