package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "open",
	Short: "open files and directories",
	Long:  "https://keith.github.io/xcode-manpages/open.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("app", "a", "", "Open the file with the specified application")
	rootCmd.Flags().StringP("bundle", "b", "", "Open the file with the specified bundle identifier")
	rootCmd.Flags().BoolP("fork", "f", false, "Open the file with the default application as a new fork")
	rootCmd.Flags().BoolP("fresh", "F", false, "Launch the application fresh, that is, without restoring windows")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("hide", "g", false, "Do not bring the application to the foreground")
	rootCmd.Flags().BoolP("j", "j", false, "Launch the app hidden")
	rootCmd.Flags().BoolP("new", "n", false, "Open a new instance of the application even if one is already running")
	rootCmd.Flags().BoolP("reveal", "R", false, "Reveal the file in the Finder instead of opening it")
	rootCmd.Flags().StringP("sdk", "s", "", "Choose an SDK when opening an application built for multiple platforms")
	rootCmd.Flags().BoolP("text", "e", false, "Open the file with /Applications/TextEdit.app")
	rootCmd.Flags().BoolP("text-editor", "t", false, "Open the file with the default text editor")
	rootCmd.Flags().StringP("url", "u", "", "Open the specified URL")
	rootCmd.Flags().BoolP("wait", "W", false, "Wait for the opened application to exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"app":    carapace.ActionFiles(".app"),
		"bundle": carapace.ActionValues(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
