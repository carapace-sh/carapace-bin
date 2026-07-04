package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "attrib",
	Short: "display, set, or clear the read-only, archive, system, and hidden attributes",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/attrib",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("s", "s", false, "process files in the current directory and all subdirectories")
	rootCmd.Flags().BoolP("d", "d", false, "process directories as well")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
