package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Read a configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_configCmd).Standalone()

	repository_helpCmd.AddCommand(repository_help_configCmd)
}
