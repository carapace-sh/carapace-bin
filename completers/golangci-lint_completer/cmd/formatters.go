package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/golangci-lint_completer/cmd/action"
	"github.com/spf13/cobra"
)

var formattersCmd = &cobra.Command{
	Use:   "formatters",
	Short: "List current formatters configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formattersCmd).Standalone()

	formattersCmd.Flags().StringP("config", "c", "", "Read config from file path `PATH`")
	formattersCmd.Flags().String("default", "", "Default set of linters to enable")
	formattersCmd.Flags().StringSliceP("disable", "D", nil, "Disable specific linter")
	formattersCmd.Flags().StringSliceP("enable", "E", nil, "Enable specific linter")
	formattersCmd.Flags().StringSlice("enable-only", nil, "Override linters configuration section to only run the specific linter(s)")
	formattersCmd.Flags().Bool("fast-only", false, "Filter enabled linters to run only fast linters")
	formattersCmd.Flags().Bool("json", false, "Display as JSON")
	formattersCmd.Flags().Bool("no-config", false, "Don't read config file")
	rootCmd.AddCommand(formattersCmd)

	carapace.Gen(formattersCmd).FlagCompletion(carapace.ActionMap{
		"config":      carapace.ActionFiles(),
		"default":     carapace.ActionValues("standard"), // TODO
		"disable":     action.ActionLinters(formattersCmd).UniqueList(","),
		"enable":      action.ActionLinters(formattersCmd).UniqueList(","),
		"enable-only": action.ActionLinters(formattersCmd).UniqueList(","),
	})
}
