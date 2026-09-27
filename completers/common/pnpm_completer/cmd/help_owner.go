package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ownerCmd = &cobra.Command{
	Use:   "owner",
	Short: "Manage package owners on the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ownerCmd).Standalone()

	helpCmd.AddCommand(help_ownerCmd)
}
