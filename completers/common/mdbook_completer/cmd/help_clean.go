package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Deletes a built book",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_cleanCmd).Standalone()

	helpCmd.AddCommand(help_cleanCmd)
}
