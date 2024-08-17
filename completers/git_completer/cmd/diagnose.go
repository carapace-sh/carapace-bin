package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diagnoseCmd = &cobra.Command{
	Use:     "diagnose",
	Short:   "Generate a zip archive of diagnostic information",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(diagnoseCmd).Standalone()

	diagnoseCmd.Flags().String("mode", "", "specify the type of diagnostics that should be collected")
	diagnoseCmd.Flags().StringP("output-directory", "o", "", "place the resulting diagnostics archive in <path>")
	diagnoseCmd.Flags().StringP("suffix", "s", "", "specify an alternate suffix for the diagnostics archive name")
	rootCmd.AddCommand(diagnoseCmd)

	carapace.Gen(diagnoseCmd).FlagCompletion(carapace.ActionMap{
		"mode":             carapace.ActionValues("stats", "all"),
		"output-directory": carapace.ActionDirectories(),
	})
}
