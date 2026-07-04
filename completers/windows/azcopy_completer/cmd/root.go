package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "azcopy",
	Short: "command-line utility for copying data to/from Azure Storage",
	Long:  "https://learn.microsoft.com/en-us/azure/storage/common/storage-use-azcopy-v10",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"copy", "copy data from source to destination",
			"sync", "synchronize a directory to another",
			"remove", "remove files or directories",
			"list", "list entities in a resource",
			"login", "login to Azure Active Directory",
			"logout", "logout from Azure Active Directory",
			"env", "show environment variables",
			"jobs", "manage jobs",
			"benchmark", "run a performance benchmark",
		),
	)
}
