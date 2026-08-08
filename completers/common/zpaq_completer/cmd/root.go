package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "zpaq",
	Short:              "journaling archiver for incremental backups",
	Long:               "https://mattmahoney.net/dc/zpaq.html",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"add", "append changes in files to archive",
			"a", "append changes in files to archive",
			"extract", "extract files from archive",
			"x", "extract files from archive",
			"list", "list the archive contents",
			"l", "list the archive contents",
		),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}

			switch c.Args[0] {
			case "add", "a":
				return carapace.ActionExecute(addCmd).Shift(1)
			case "extract", "x":
				return carapace.ActionExecute(extractCmd).Shift(1)
			case "list", "l":
				return carapace.ActionExecute(listCmd).Shift(1)
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
