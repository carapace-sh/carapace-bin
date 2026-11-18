package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Work with captured command output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_outputCmd).Standalone()

	helpCmd.AddCommand(help_outputCmd)
}
