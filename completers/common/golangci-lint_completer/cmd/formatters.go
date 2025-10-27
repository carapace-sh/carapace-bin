package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/golangci-lint_completer/cmd/action"
	"github.com/spf13/cobra"
)

var formattersCmd = &cobra.Command{
	Use:   "formatters",
	Short: "List current formatters configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formattersCmd).Standalone()

	formattersCmd.Flags().StringP("config", "c", "", "Read config from file path `PATH`")
	formattersCmd.Flags().StringSliceP("enable", "E", nil, "Enable specific formatter")
	formattersCmd.Flags().Bool("json", false, "Display as JSON")
	formattersCmd.Flags().Bool("no-config", false, "Don't read config file")
	rootCmd.AddCommand(formattersCmd)

	carapace.Gen(formattersCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"enable": action.ActionLinters(formattersCmd).UniqueList(","),
	})
}
