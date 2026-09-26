package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var self_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display uv's version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(self_versionCmd).Standalone()

	self_versionCmd.Flags().String("output-format", "text", "")
	self_versionCmd.Flags().Bool("short", false, "Only print the version")
	selfCmd.AddCommand(self_versionCmd)
	carapace.Gen(self_versionCmd).FlagCompletion(carapace.ActionMap{
		"output-format": carapace.ActionValuesDescribed(
			"text", "Display the version as plain text",
			"json", "Display the version as JSON",
		),
	})
}
