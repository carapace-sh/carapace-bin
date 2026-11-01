package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a package to the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_publishCmd).Standalone()

	helpCmd.AddCommand(help_publishCmd)
}
