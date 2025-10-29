package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_selfCmd = &cobra.Command{
	Use:   "self",
	Short: "Modify the rustup installation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_selfCmd).Standalone()

	helpCmd.AddCommand(help_selfCmd)
}
