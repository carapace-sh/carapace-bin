package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var monitoringCmd = &cobra.Command{
	Use:   "monitoring",
	Short: "[Beta] Display commands to manage monitoring",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(monitoringCmd).Standalone()
	rootCmd.AddCommand(monitoringCmd)
}
