package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_contributorsCmd = &cobra.Command{
	Use:   "contributors",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_contributorsCmd).Standalone()

	helpCmd.AddCommand(help_contributorsCmd)
}
