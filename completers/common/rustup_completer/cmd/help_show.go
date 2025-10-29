package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the active and installed toolchains or profiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_showCmd).Standalone()

	helpCmd.AddCommand(help_showCmd)
}
