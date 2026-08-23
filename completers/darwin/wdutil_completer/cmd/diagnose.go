package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "run diagnostic tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagnoseCmd).Standalone()
	rootCmd.AddCommand(diagnoseCmd)

	diagnoseCmd.Flags().StringP("output", "f", "", "Output directory path")
	diagnoseCmd.Flags().BoolP("quiet", "q", false, "Suppress legal prompt and Finder window")

	carapace.Gen(diagnoseCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionDirectories(),
	})
}
