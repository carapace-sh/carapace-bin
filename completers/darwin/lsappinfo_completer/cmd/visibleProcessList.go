package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var visibleProcessListCmd = &cobra.Command{
	Use:   "visibleProcessList",
	Short: "Show the visible (front-to-back) application list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(visibleProcessListCmd).Standalone()
	rootCmd.AddCommand(visibleProcessListCmd)
}
