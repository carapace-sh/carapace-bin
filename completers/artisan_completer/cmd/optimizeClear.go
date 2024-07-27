package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var optimizeClearCmd = &cobra.Command{
	Use:   "optimize:clear",
	Short: "Remove the cached bootstrap files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(optimizeClearCmd).Standalone()

	rootCmd.AddCommand(optimizeClearCmd)
}
