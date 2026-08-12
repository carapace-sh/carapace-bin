package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Read a configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_configCmd).Standalone()

	help_repositoryCmd.AddCommand(help_repository_configCmd)
}
