package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "Set the default toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_defaultCmd).Standalone()

	helpCmd.AddCommand(help_defaultCmd)
}
