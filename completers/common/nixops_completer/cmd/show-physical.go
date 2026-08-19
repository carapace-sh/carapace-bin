package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ShowPhysicalCmd = &cobra.Command{
	Use:   "show-physical",
	Short: "Show Physical",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ShowPhysicalCmd).Standalone()
	rootCmd.AddCommand(ShowPhysicalCmd)
}
