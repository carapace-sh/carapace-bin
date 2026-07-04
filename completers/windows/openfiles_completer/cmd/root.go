package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "openfiles",
	Short: "query, display, or disconnect files and directories opened by remote users",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/openfiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"query", "display opened files or directories",
			"disconnect", "disconnect opened files or directories",
			"local", "enable/disable local file list display",
		),
	)
}
