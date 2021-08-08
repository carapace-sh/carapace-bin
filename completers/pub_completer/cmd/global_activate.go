package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var global_activateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Make a package's executables globally available",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_activateCmd).Standalone()

	global_activateCmd.Flags().BoolP("executable", "x", false, "Executable(s) to place on PATH.")
	global_activateCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	global_activateCmd.Flags().StringP("hosted-url", "u", "", "A custom pub server URL for the package.")
	global_activateCmd.Flags().Bool("no-executables", false, "Do not put executables on PATH.")
	global_activateCmd.Flags().Bool("overwrite", false, "Overwrite executables from other packages with the same name.")
	global_activateCmd.Flags().StringP("source", "s", "", "The source used to find the package.")
	globalCmd.AddCommand(global_activateCmd)

	carapace.Gen(global_activateCmd).FlagCompletion(carapace.ActionMap{
		"source": carapace.ActionValues("git", "hosted", "path"),
	})
}
