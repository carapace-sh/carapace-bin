package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register with the configured server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_registerCmd).Standalone()

	helpCmd.AddCommand(help_registerCmd)
}
