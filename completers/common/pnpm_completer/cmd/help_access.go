package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_accessCmd = &cobra.Command{
	Use:   "access",
	Short: "Manage package access and visibility on the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_accessCmd).Standalone()

	helpCmd.AddCommand(help_accessCmd)
}
