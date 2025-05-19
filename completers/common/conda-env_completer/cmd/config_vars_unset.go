package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/conda"
	"github.com/spf13/cobra"
)

var config_vars_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset environment variables for a conda environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_vars_unsetCmd).Standalone()

	config_vars_unsetCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	config_vars_unsetCmd.Flags().StringP("name", "n", "", "Name of environment")
	config_vars_unsetCmd.Flags().StringP("prefix", "p", "", "Full path to environment location")
	config_varsCmd.AddCommand(config_vars_unsetCmd)

	carapace.Gen(config_vars_unsetCmd).FlagCompletion(carapace.ActionMap{
		"name": conda.ActionEnvironments(),
	})

	carapace.Gen(config_vars_unsetCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return conda.ActionEnvironmentVariables(config_vars_unsetCmd.Flag("name").Value.String())
		}),
	)
}
