package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showPhysicalCmd = &cobra.Command{
	Use:   "show-physical",
	Short: "print the physical network expression",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showPhysicalCmd).Standalone()
	rootCmd.AddCommand(showPhysicalCmd)
}
