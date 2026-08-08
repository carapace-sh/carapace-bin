package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var gpgKey_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of GPG keys for the currently authenticated user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gpgKey_listCmd).Standalone()

	gpgKey_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	gpgKey_listCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	gpgKey_listCmd.Flags().Bool("show-id", false, "Shows IDs of GPG keys.")
	gpgKeyCmd.AddCommand(gpgKey_listCmd)

	carapace.Gen(gpgKey_getCmd).FlagCompletion(carapace.ActionMap{
		"jq":     jq.ActionFilters(),
		"output": carapace.ActionValues("text", "json"),
	})
}
