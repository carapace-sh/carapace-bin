package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bugreportCmd = &cobra.Command{
	Use:     "bugreport",
	Short:   "Collect information for user to file a bug report",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(bugreportCmd).Standalone()

	bugreportCmd.Flags().String("diagnose", "", "Create a zip archive of supplemental information about the user’s machine")
	bugreportCmd.Flags().Bool("no-diagnose", false, "Create a zip archive of supplemental information about the user’s machine")
	bugreportCmd.Flags().StringP("output-directory", "o", "", "Place the resulting bug report file in <path> instead of the current directory")
	bugreportCmd.Flags().StringP("suffix", "s", "", "Specify an alternate suffix for the bugreport name")
	rootCmd.AddCommand(bugreportCmd)

	bugreportCmd.Flag("diagnose").NoOptDefVal = " "

	carapace.Gen(bugreportCmd).FlagCompletion(carapace.ActionMap{
		"diagnose":         carapace.ActionValues("stats", "all"),
		"output-directory": carapace.ActionDirectories(),
	})
}
