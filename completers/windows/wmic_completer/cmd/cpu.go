package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "CPU management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cpuCmd).Standalone()
	rootCmd.AddCommand(cpuCmd)
}
