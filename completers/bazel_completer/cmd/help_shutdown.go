package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_shutdownCmd).Standalone()

	helpCmd.AddCommand(help_shutdownCmd)
}
