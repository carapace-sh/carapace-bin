package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/conda"
	"github.com/spf13/cobra"
)

var config_vars_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set environment variables for a conda environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_vars_setCmd).Standalone()

	config_vars_setCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	config_vars_setCmd.Flags().StringP("name", "n", "", "Name of environment")
	config_vars_setCmd.Flags().StringP("prefix", "p", "", "Full path to environment location")
	config_varsCmd.AddCommand(config_vars_setCmd)

	carapace.Gen(config_vars_setCmd).FlagCompletion(carapace.ActionMap{
		"name": conda.ActionEnvironments(),
	})

	carapace.Gen(config_vars_setCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return conda.ActionEnvironmentVariables(config_vars_setCmd.Flag("name").Value.String()).Suffix("=")
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
