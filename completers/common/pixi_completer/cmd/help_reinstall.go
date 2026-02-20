package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_reinstallCmd = &cobra.Command{
	Use:   "reinstall",
	Short: "Re-install an environment, both updating the lockfile and re-installing the environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_reinstallCmd).Standalone()

	helpCmd.AddCommand(help_reinstallCmd)
}
