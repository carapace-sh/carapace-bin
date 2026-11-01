package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs an arbitrary command specified in the package's start property of its scripts object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_startCmd).Standalone()

	helpCmd.AddCommand(help_startCmd)
}
