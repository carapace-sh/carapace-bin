package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Stage changes. With no paths, stages everything",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_addCmd).Standalone()

	helpCmd.AddCommand(help_addCmd)
}
