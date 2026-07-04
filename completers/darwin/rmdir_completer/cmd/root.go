package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rmdir",
	Short: "remove directories",
	Long:  "https://keith.github.io/xcode-manpages/rmdir.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("parents", "p", false, "Remove the directory and any missing parent directories")
	rootCmd.Flags().BoolP("verbose", "v", false, "Cause rmdir to be verbose, showing directories as they are removed")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionDirectories())
}
