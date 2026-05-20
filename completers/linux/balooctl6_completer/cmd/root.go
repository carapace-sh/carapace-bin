package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "balooctl6",
	Short: "CLI for the KDE Baloo file indexing framework",
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
		"format":        carapace.ActionValues("multiline", "json", "simple").StyleF(style.ForKeyword),
		"qmljsdebugger": carapace.ActionValues("port:1234", "port:1234,block"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"status", "Print the status of the indexer",
			"enable", "Enable the file indexer",
			"disable", "Disable the file indexer",
			"purge", "Remove the index database",
			"suspend", "Suspend the file indexer",
			"resume", "Resume the file indexer",
			"check", "Check for any unindexed files and index them",
			"index", "Index the specified files",
			"clear", "Forget the specified files",
			"config", "Modify the Baloo configuration",
			"monitor", "Monitor the file indexer",
			"indexSize", "Display the disk space used by index",
			"failed", "Display files which could not be indexed",
		),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}
			switch c.Args[0] {
			case "status", "index", "clear":
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
