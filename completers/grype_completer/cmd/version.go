package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().StringP("output", "o", "text", "format to display results")
	rootCmd.AddCommand(versionCmd)

	carapace.Gen(versionCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
	})
}
