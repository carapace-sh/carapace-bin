package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_starCmd = &cobra.Command{
	Use:   "star",
	Short: "Marks a package as a favorite",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_starCmd).Standalone()

	helpCmd.AddCommand(help_starCmd)
}
