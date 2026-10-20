package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Describe each key in a mise TOML file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_explainCmd).Standalone()

	config_explainCmd.Flags().StringP("file", "f", "", "Config file to read, or a directory whose config file to read")
	config_explainCmd.Flags().BoolP("global", "g", false, "Read the global config file")
	config_explainCmd.Flags().BoolP("help", "h", false, "Print help")
	config_explainCmd.Flags().BoolP("json", "J", false, "Print the keys as JSON")
	config_explainCmd.Flags().Bool("override", false, "Read the global override file, `config.override.toml`")
	config_explainCmd.Flags().String("path", "", "Config file to read, or a directory whose config file to read")
	config_explainCmd.Flags().Bool("system", false, "Read the system config file")
	configCmd.AddCommand(config_explainCmd)

	carapace.Gen(config_explainCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(config_explainCmd).PositionalCompletion(
		action.ConfigKeys("explain"),
	)
}
