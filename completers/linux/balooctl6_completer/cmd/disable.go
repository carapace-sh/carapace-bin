package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable the file indexer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disableCmd).Standalone()

	rootCmd.AddCommand(disableCmd)
}
