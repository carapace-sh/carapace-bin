package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mise",
	Short: "Polyglot runtime manager and task runner",
	Long:  "https://mise.jdx.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("cd", "", "Change to this directory before running mise")
	rootCmd.PersistentFlags().StringP("config", "c", "", "Path to mise config file")
	rootCmd.PersistentFlags().StringP("env", "E", "", "Set environment")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print help")
	rootCmd.PersistentFlags().String("log-level", "", "Set the log level")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Suppress non-error output")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Show extra output")
	rootCmd.PersistentFlags().BoolP("version", "V", false, "Print version")
	rootCmd.Flags().Bool("yes", false, "Automatically answer yes to prompts")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cd":        carapace.ActionDirectories(),
		"config":    carapace.ActionFiles(".toml"),
		"log-level": carapace.ActionValues("trace", "debug", "info", "warn", "error"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		mise.ActionTasks().FilterArgs(),
	)
}
