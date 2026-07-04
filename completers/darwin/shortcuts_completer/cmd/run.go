package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute a shortcut",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().StringSliceP("input-path", "i", nil, "Input files to pass to the shortcut")
	runCmd.Flags().StringP("output-path", "o", "", "Write the shortcut's output to a file")
	runCmd.Flags().String("output-type", "", "Force output to a specific UTI")
	rootCmd.AddCommand(runCmd)
	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"input-path":  carapace.ActionFiles(),
		"output-path": carapace.ActionFiles(),
	})
}
