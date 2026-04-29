package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "balooctl6",
	Short: "KDE Baloo file indexer control tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error { return rootCmd.Execute() }

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("format", "f", "", "output format")
	rootCmd.Flags().BoolP("version", "v", false, "display version information")
	rootCmd.Flags().BoolP("help", "h", false, "display help")
	rootCmd.Flags().Bool("help-all", false, "display help including generic Qt options")
	rootCmd.Flags().String("qmljsdebugger", "", "activate QML/JS debugger with port specification")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"format":        carapace.ActionValues("multiline", "json", "simple").StyleF(style.ForKeyword),
		"qmljsdebugger": carapace.ActionValues("port:1234", "port:1234,block"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"status", "print the status of the indexer",
			"enable", "enable the file indexer",
			"disable", "disable the file indexer",
			"purge", "remove the index database",
			"suspend", "suspend the file indexer",
			"resume", "resume the file indexer",
			"check", "check for unindexed files and index them",
			"index", "index the specified files",
			"clear", "forget the specified files",
			"config", "modify the Baloo configuration",
			"monitor", "monitor the file indexer",
			"indexSize", "display disk space used by the index",
			"failed", "display files which could not be indexed",
		).StyleF(style.ForKeyword),
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
