package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_whichCmd = &cobra.Command{
	Use:   "which",
	Short: "Display which binary will be run for a given command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_whichCmd).Standalone()

	helpCmd.AddCommand(help_whichCmd)
}
