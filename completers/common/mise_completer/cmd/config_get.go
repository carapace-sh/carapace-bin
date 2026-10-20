package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Print a value from one mise TOML file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCmd).Standalone()

	config_getCmd.Flags().StringP("file", "f", "", "Config file to read, or a directory whose config file to read")
	config_getCmd.Flags().BoolP("global", "g", false, "Read the global config file")
	config_getCmd.Flags().BoolP("help", "h", false, "Print help")
	config_getCmd.Flags().Bool("override", false, "Read the global override file, `config.override.toml`")
	config_getCmd.Flags().String("path", "", "Config file to read, or a directory whose config file to read")
	config_getCmd.Flags().Bool("system", false, "Read the system config file")
	configCmd.AddCommand(config_getCmd)

	carapace.Gen(config_getCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(config_getCmd).PositionalCompletion(
		action.ConfigKeys("get"),
	)
}
