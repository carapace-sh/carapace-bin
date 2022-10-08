package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Display resource (CPU/memory) usage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(topCmd).Standalone()
	rootCmd.AddCommand(topCmd)
}
