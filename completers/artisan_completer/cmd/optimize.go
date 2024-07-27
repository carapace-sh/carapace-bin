package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var optimizeCmd = &cobra.Command{
	Use:   "optimize",
	Short: "Cache framework bootstrap, configuration, and metadata to increase performance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(optimizeCmd).Standalone()

	rootCmd.AddCommand(optimizeCmd)
}
