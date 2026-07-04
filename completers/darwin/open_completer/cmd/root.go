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
	rootCmd.Flags().StringP("args", "args", "", "Pass remaining arguments as arguments to the application")
	rootCmd.Flags().StringP("bundle", "b", "", "Open the file with the specified bundle identifier")
	rootCmd.Flags().StringP("env", "E", "", "Set the environment variable when launching the application")
	rootCmd.Flags().BoolP("fresh", "F", false, "Launch the application fresh")
	rootCmd.Flags().BoolP("hide", "g", false, "Do not bring the application to the foreground")
	rootCmd.Flags().BoolP("new", "n", false, "Open a new instance of the application")
	rootCmd.Flags().BoolP("print", "p", false, "Open the file with the default printing application")
	rootCmd.Flags().BoolP("reveal", "R", false, "Reveal the file in the Finder")
	rootCmd.Flags().StringP("stderr", "e", "", "Redirect stderr to the specified file")
	rootCmd.Flags().StringP("stdin", "s", "", "Redirect stdin from the specified file")
	rootCmd.Flags().StringP("stdout", "o", "", "Redirect stdout to the specified file")
	rootCmd.Flags().BoolP("url", "u", false, "Treat the argument as a URL")
	rootCmd.Flags().StringP("with", "W", "", "Open the file with the specified application")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
