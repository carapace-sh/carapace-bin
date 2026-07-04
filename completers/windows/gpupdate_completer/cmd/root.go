package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gpupdate",
	Short: "update group policy settings",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/gpupdate",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("force", false, "reapply all policy settings")
}
