package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_blameCmd = &cobra.Command{
	Use:   "blame",
	Short: "Show who last modified each line of a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_blameCmd).Standalone()

	helpCmd.AddCommand(help_blameCmd)
}
