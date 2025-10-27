package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/golangci-lint_completer/cmd/action"
	"github.com/spf13/cobra"
)

var lintersCmd = &cobra.Command{
	Use:   "linters",
	Short: "List current linters configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lintersCmd).Standalone()

	lintersCmd.Flags().StringP("config", "c", "", "Read config from file path `PATH`")
	lintersCmd.Flags().String("default", "", "Default set of linters to enable")
	lintersCmd.Flags().StringSliceP("disable", "D", nil, "Disable specific linter")
	lintersCmd.Flags().StringSliceP("enable", "E", nil, "Enable specific linter")
	lintersCmd.Flags().StringSlice("enable-only", nil, "Override linters configuration section to only run the specific linter(s)")
	lintersCmd.Flags().Bool("fast-only", false, "Filter enabled linters to run only fast linters")
	lintersCmd.Flags().Bool("json", false, "Display as JSON")
	lintersCmd.Flags().Bool("no-config", false, "Don't read config file")
	rootCmd.AddCommand(lintersCmd)

	carapace.Gen(lintersCmd).FlagCompletion(carapace.ActionMap{
		"config":      carapace.ActionFiles(),
		"default":     carapace.ActionValues("standard"), // TODO
		"disable":     action.ActionLinters(lintersCmd).UniqueList(","),
		"enable":      action.ActionLinters(lintersCmd).UniqueList(","),
		"enable-only": action.ActionLinters(lintersCmd).UniqueList(","),
	})
}
