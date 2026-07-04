package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cd",
	Short: "display or change the current directory",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/cd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("d", "d", false, "change current drive in addition to current directory")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}
