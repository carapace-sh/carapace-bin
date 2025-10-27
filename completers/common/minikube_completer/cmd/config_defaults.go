package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_defaultsCmd = &cobra.Command{
	Use:   "defaults PROPERTY_NAME",
	Short: "Lists all valid default values for PROPERTY_NAME",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_defaultsCmd).Standalone()

	config_defaultsCmd.Flags().StringP("output", "o", "", "Output format. Accepted values: [json, yaml]")
	configCmd.AddCommand(config_defaultsCmd)

	carapace.Gen(config_defaultsCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("json", "yaml"),
	})

	carapace.Gen(config_defaultsCmd).PositionalCompletion(
		action.ActionConfigNames(),
	)
}
