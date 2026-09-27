package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a package from a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_deployCmd).Standalone()

	helpCmd.AddCommand(help_deployCmd)
}
