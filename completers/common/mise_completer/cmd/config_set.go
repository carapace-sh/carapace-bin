package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value in one mise TOML file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()

	config_setCmd.Flags().Bool("append", false, "Append the value to a list without duplicating an existing entry")
	config_setCmd.Flags().StringP("file", "f", "", "Config file to edit, or a directory whose config file to edit")
	config_setCmd.Flags().BoolP("global", "g", false, "Edit the global config file")
	config_setCmd.Flags().BoolP("help", "h", false, "Print help")
	config_setCmd.Flags().Bool("override", false, "Edit the global override file, `config.override.toml`")
	config_setCmd.Flags().String("path", "", "Config file to edit, or a directory whose config file to edit")
	config_setCmd.Flags().Bool("remove", false, "Remove the value from an existing list")
	config_setCmd.Flags().Bool("system", false, "Edit the system config file")
	config_setCmd.Flags().StringP("type", "t", "infer", "TOML type to store the value as; inferred from the value by default")
	configCmd.AddCommand(config_setCmd)

	carapace.Gen(config_setCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
		"type": carapace.ActionValues("bool", "float", "infer", "integer", "list", "set", "string"),
	})

	carapace.Gen(config_setCmd).PositionalCompletion(
		action.ActionConfigKeys("set"),
	)
}
