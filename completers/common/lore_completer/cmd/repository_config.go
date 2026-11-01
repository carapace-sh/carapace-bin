package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Read a configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_configCmd).Standalone()

	repository_configCmd.Flags().BoolP("help", "h", false, "Print help")
	repositoryCmd.AddCommand(repository_configCmd)
}
