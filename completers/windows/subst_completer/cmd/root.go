package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "subst",
	Short: "associate a path with a drive letter",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/subst",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("d", "d", false, "delete a substituted drive letter")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionDirectories(),
	)
}
