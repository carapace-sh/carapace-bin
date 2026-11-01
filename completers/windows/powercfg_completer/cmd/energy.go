package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var energyCmd = &cobra.Command{
	Use:   "energy",
	Short: "analyze the system for energy efficiency problems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(energyCmd).Standalone()
	rootCmd.AddCommand(energyCmd)
}
