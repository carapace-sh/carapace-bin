package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ciCmd = &cobra.Command{
	Use:   "ci",
	Short: "Runs clean then install with a frozen lockfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ciCmd).Standalone()

	helpCmd.AddCommand(help_ciCmd)
}
