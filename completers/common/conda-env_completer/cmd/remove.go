package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/conda"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an environmentRemoves a provided environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().BoolP("dry-run", "d", false, "Only display what would have been done")
	removeCmd.Flags().BoolP("help", "h", false, "Show this help message and exit")
	removeCmd.Flags().Bool("json", false, "Report all output as json")
	removeCmd.Flags().StringP("name", "n", "", "Name of environment")
	removeCmd.Flags().StringP("prefix", "p", "", "Full path to environment location")
	removeCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar.")
	removeCmd.Flags().String("solver", "", "Choose which solver backend to use")
	removeCmd.Flags().BoolP("verbose", "v", false, "Once for INFO, twice for DEBUG, three times for TRACE.")
	removeCmd.Flags().BoolP("yes", "y", false, "Sets any confirmation values to 'yes' automatically")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"name":   conda.ActionEnvironments(),
		"solver": carapace.ActionValues("classic"),
	})
}
