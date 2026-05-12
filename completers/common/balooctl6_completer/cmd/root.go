package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "balooctl6",
	Short: "control the Baloo file indexer",
	Long:  "https://community.kde.org/Baloo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("format", "f", "multiline", "Output format")
	rootCmd.PersistentFlags().Bool("help-all", false, "Displays help, including generic Qt options")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Displays help on commandline options")
	rootCmd.PersistentFlags().String("qmljsdebugger", "", "Activates the QML/JS debugger with a specified port")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Displays version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValuesDescribed(
			"multiline", "multi-line file status output",
			"json", "JSON file status output",
			"simple", "single-line file status output",
		),
	})

	rootCmd.AddCommand(
		newConfigCommand(),
		newFileCommand("clear", "Forget the specified files"),
		newCommand("check", "Check for any unindexed files and index them"),
		newCommand("disable", "Disable the file indexer"),
		newCommand("enable", "Enable the file indexer"),
		newCommand("failed", "Display files which could not be indexed"),
		newFileCommand("index", "Index the specified files"),
		newCommand("indexSize", "Display the disk space used by index"),
		newCommand("monitor", "Monitor the file indexer"),
		newCommand("purge", "Remove the index database"),
		newCommand("resume", "Resume the file indexer"),
		newFileCommand("status", "Print the status of the indexer"),
		newCommand("suspend", "Suspend the file indexer"),
	)
}

func newCommand(use string, short string) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
}

func newFileCommand(use string, short string) *cobra.Command {
	command := newCommand(use, short)

	carapace.Gen(command).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)

	return command
}
