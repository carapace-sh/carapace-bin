package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/conda"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the current environment based on environment file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().StringP("file", "f", "", "environment definition")
	updateCmd.Flags().BoolP("help", "h", false, "Show this help message and exit")
	updateCmd.Flags().Bool("json", false, "Report all output as json")
	updateCmd.Flags().StringP("name", "n", "", "Name of environment")
	updateCmd.Flags().StringP("prefix", "p", "", "Full path to environment location")
	updateCmd.Flags().Bool("prune", false, "remove installed packages not defined in environment.yml")
	updateCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar")
	updateCmd.Flags().String("solver", "", "Choose which solver backend to use")
	updateCmd.Flags().BoolP("verbose", "v", false, "Use once for info, twice for debug, three times for trace")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"file":   carapace.ActionFiles(),
		"name":   conda.ActionEnvironments(),
		"solver": carapace.ActionValues("classic"),
	})
}
