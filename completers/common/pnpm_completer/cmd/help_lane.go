package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_laneCmd = &cobra.Command{
	Use:   "lane",
	Short: "Manage per-package release lanes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_laneCmd).Standalone()

	helpCmd.AddCommand(help_laneCmd)
}
