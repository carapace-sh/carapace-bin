package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_llCmd = &cobra.Command{
	Use:   "ll",
	Short: "List installed packages in long format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_llCmd).Standalone()

	helpCmd.AddCommand(help_llCmd)
}
