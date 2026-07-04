package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "touch",
	Short: "change file access and modification times",
	Long:  "https://keith.github.io/xcode-manpages/touch.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "Change the access time of the file")
	rootCmd.Flags().BoolS("c", "c", false, "Do not create the file if it does not exist")
	rootCmd.Flags().StringS("d", "d", "", "Use the specified time instead of the current time")
	rootCmd.Flags().BoolS("f", "f", false, "Attempt to force the update, even if the file permissions do not allow it")
	rootCmd.Flags().BoolS("h", "h", false, "If the file is a symbolic link, change the times of the link itself")
	rootCmd.Flags().BoolS("m", "m", false, "Change the modification time of the file")
	rootCmd.Flags().StringP("reference", "r", "", "Use the access and modifications times from the specified file")
	rootCmd.Flags().StringS("t", "t", "", "Use the specified time instead of the current time")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"reference": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
