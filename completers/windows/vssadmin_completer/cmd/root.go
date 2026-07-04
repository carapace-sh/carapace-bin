package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vssadmin",
	Short: "volume shadow copy service administration",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/vssadmin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"list", "list shadow copies, writers, providers, or volumes",
			"delete", "delete shadow copies",
			"resize", "resize shadow copy storage",
			"revert", "revert a shadow copy",
			"add", "add a shadow copy",
		),
	)
}
