package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/composer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "composer",
	Short: "Dependency Manager for PHP",
	Long:  "https://getcomposer.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("ansi", false, "Force (or disable --no-ansi) ANSI output")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Display help for the given command. When no command is given display help for the list command")
	rootCmd.PersistentFlags().Bool("no-ansi", false, "Negate the \"--ansi\" option")
	rootCmd.PersistentFlags().Bool("no-cache", false, "Prevent use of the cache")
	rootCmd.PersistentFlags().BoolP("no-interaction", "n", false, "Do not ask any interactive question")
	rootCmd.PersistentFlags().Bool("no-plugins", false, "Whether to disable plugins.")
	rootCmd.PersistentFlags().Bool("no-scripts", false, "Skips the execution of all scripts defined in composer.json file.")
	rootCmd.PersistentFlags().Bool("profile", false, "Display timing and memory usage information")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Do not output any message")
	rootCmd.PersistentFlags().CountP("verbose", "v", "Increase the verbosity of messages: 1 for normal output, 2 for more verbose output and 3 for debug")
	rootCmd.PersistentFlags().BoolP("version", "V", false, "Display this application version")
	rootCmd.PersistentFlags().StringP("working-dir", "d", "", "If specified, use the given directory as working directory.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"working-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		action.ActionScripts(rootCmd),
	)
}
