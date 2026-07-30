package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "popd",
	Short: "change the current directory to the directory stored by the pushd command",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/popd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
