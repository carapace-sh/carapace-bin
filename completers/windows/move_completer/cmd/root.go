package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "move",
	Short: "move one or more files from one directory to another directory",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/move",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("y", "y", false, "suppress prompting to confirm overwriting")
	rootCmd.Flags().BoolP("n", "n", false, "do not overwrite an existing file")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
