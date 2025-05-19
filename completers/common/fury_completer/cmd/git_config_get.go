package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_config_getCmd = &cobra.Command{
	Use:   "get REPO KEY",
	Short: "Get Git build environment key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_config_getCmd).Standalone()

	git_configCmd.AddCommand(git_config_getCmd)
}
