package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "expand",
	Short: "expand one or more compressed files from a distribution cabinet (.cab) file",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/expand",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("f", "f", false, "specify files to expand from a cabinet")
	rootCmd.Flags().BoolP("r", "r", false, "rename expanded files")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
