package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_unstarCmd = &cobra.Command{
	Use:   "unstar",
	Short: "Unmarks a package as a favorite",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_unstarCmd).Standalone()

	helpCmd.AddCommand(help_unstarCmd)
}
