package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var OptimizeImageCmd = &cobra.Command{
	Use:   "Optimize-Image",
	Short: "optimize an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(OptimizeImageCmd).Standalone()
	rootCmd.AddCommand(OptimizeImageCmd)
}
