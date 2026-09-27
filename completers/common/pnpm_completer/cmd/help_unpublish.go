package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_unpublishCmd = &cobra.Command{
	Use:   "unpublish",
	Short: "Removes a package from the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_unpublishCmd).Standalone()

	helpCmd.AddCommand(help_unpublishCmd)
}
