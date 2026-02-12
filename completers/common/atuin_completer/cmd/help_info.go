package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Information about dotfiles locations and ENV vars",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_infoCmd).Standalone()

	helpCmd.AddCommand(help_infoCmd)
}
