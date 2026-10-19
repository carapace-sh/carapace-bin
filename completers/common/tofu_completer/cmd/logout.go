package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout [options] <hostname>",
	Short: "Remove locally-stored credentials for a remote host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logoutCmd).Standalone()

	logoutCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	logoutCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	logoutCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	logoutCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	logoutCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	logoutCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	logoutCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	logoutCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	rootCmd.AddCommand(logoutCmd)

	carapace.Gen(logoutCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"json-into":   carapace.ActionFiles(),
	})

	carapace.Gen(logoutCmd).PositionalCompletion(
		net.ActionHosts(), // TODO complete authorized hosts
	)
}
