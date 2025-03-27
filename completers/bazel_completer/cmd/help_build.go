package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_buildCmd = &cobra.Command{
	Use:   "build",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_buildCmd).Standalone()

	helpCmd.AddCommand(help_buildCmd)
}
