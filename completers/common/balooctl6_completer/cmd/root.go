package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "balooctl6",
	Short: "CLI for KDE's Baloo file indexing and search framework",
	Long:  "https://community.kde.org/Baloo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("format", "f", "", "Output format")
	rootCmd.Flags().BoolP("help", "h", false, "Displays help on commandline options")
	rootCmd.Flags().Bool("help-all", false, "Displays help, including generic Qt options")
	rootCmd.Flags().String("qmljsdebugger", "", "Activates the QML/JS debugger with a specified port")
	rootCmd.Flags().BoolP("version", "v", false, "Displays version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("multiline", "json", "simple"),
	})

	addCommand("status", "Print the status of the indexer", carapace.ActionFiles())
	addCommand("enable", "Enable the file indexer", carapace.ActionValues())
	addCommand("disable", "Disable the file indexer", carapace.ActionValues())
	addCommand("purge", "Remove the index database", carapace.ActionValues())
	addCommand("suspend", "Suspend the file indexer", carapace.ActionValues())
	addCommand("resume", "Resume the file indexer", carapace.ActionValues())
	addCommand("check", "Check for any unindexed files and index them", carapace.ActionValues())
	addCommand("index", "Index the specified files", carapace.ActionFiles())
	addCommand("clear", "Forget the specified files", carapace.ActionFiles())
	addCommand("config", "Modify the Baloo configuration", actionConfig())
	addCommand("monitor", "Monitor the file indexer", carapace.ActionValues())
	addCommand("indexSize", "Display the disk space used by index", carapace.ActionValues())
	addCommand("failed", "Display files which could not be indexed", carapace.ActionValues())
}

func addCommand(name string, description string, completion carapace.Action) {
	cmd := &cobra.Command{
		Use:   name,
		Short: description,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	carapace.Gen(cmd).Standalone()
	rootCmd.AddCommand(cmd)

	carapace.Gen(cmd).PositionalAnyCompletion(completion)
}

func actionConfig() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch len(c.Args) {
		case 0:
			return carapace.ActionValuesDescribed(
				"add", "Add a value to a list option",
				"rm", "Remove a value from a list option",
				"remove", "Remove a value from a list option",
				"list", "List configuration values",
				"ls", "List configuration values",
				"show", "List configuration values",
				"set", "Set a scalar option",
				"help", "Show config help",
			)
		case 1:
			switch c.Args[0] {
			case "add", "rm", "remove":
				return carapace.ActionValuesDescribed(
					"includeFolders", "Folders included in indexing",
					"excludeFolders", "Folders excluded from indexing",
					"excludeFilters", "Filename filters excluded from indexing",
					"excludeMimetypes", "MIME types excluded from indexing",
				)
			case "list", "ls", "show":
				return carapace.ActionValuesDescribed(
					"hidden", "Index hidden files",
					"contentIndexing", "Index file content",
					"includeFolders", "Folders included in indexing",
					"excludeFolders", "Folders excluded from indexing",
					"excludeFilters", "Filename filters excluded from indexing",
					"excludeMimetypes", "MIME types excluded from indexing",
				)
			case "set":
				return carapace.ActionValuesDescribed(
					"hidden", "Index hidden files",
					"contentIndexing", "Index file content",
				)
			}
		case 2:
			switch c.Args[0] {
			case "add", "rm", "remove":
				switch c.Args[1] {
				case "includeFolders", "excludeFolders":
					return carapace.ActionDirectories()
				}
			case "set":
				return carapace.ActionValues("yes", "no", "true", "false", "1", "0")
			}
		}

		return carapace.ActionValues()
	})
}
