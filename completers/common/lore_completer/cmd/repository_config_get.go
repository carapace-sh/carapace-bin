package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_config_getCmd).Standalone()

	repository_config_getCmd.Flags().BoolP("help", "h", false, "Print help")
	repository_configCmd.AddCommand(repository_config_getCmd)
}
