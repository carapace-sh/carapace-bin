package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get configuration settings from the local Anchor.toml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCmd).Standalone()

	config_getCmd.Flags().BoolP("help", "h", false, "Print help")
	configCmd.AddCommand(config_getCmd)
}
