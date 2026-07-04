package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "setspn",
	Short: "read, modify, and delete Service Principal Names (SPNs)",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/setspn",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"add", "add an SPN",
			"delete", "delete an SPN",
			"find", "find duplicate SPNs",
			"list", "list SPNs for an account",
			"query", "query SPNs for an account",
		),
	)
}
