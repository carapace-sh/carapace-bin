package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_modCmd = &cobra.Command{
	Use:   "mod",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_modCmd).Standalone()

	helpCmd.AddCommand(help_modCmd)
}
