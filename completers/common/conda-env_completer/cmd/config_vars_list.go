package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/conda"
	"github.com/spf13/cobra"
)

var config_vars_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List environment variables for a conda environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_vars_listCmd).Standalone()

	config_vars_listCmd.Flags().BoolP("help", "h", false, "Show this help message and exit")
	config_vars_listCmd.Flags().Bool("json", false, "Report all output as json")
	config_vars_listCmd.Flags().StringP("name", "n", "", "Name of environment")
	config_vars_listCmd.Flags().StringP("prefix", "p", "", "Full path to environment location")
	config_vars_listCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar")
	config_vars_listCmd.Flags().BoolP("verbose", "v", false, "Use once for info, twice for debug, three times for trace")
	config_varsCmd.AddCommand(config_vars_listCmd)

	carapace.Gen(config_vars_listCmd).FlagCompletion(carapace.ActionMap{
		"name": conda.ActionEnvironments(),
	})
}
