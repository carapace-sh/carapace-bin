package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_approvals_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List project commands and their approval status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_approvals_listCmd).Standalone()

	hook_approvals_listCmd.Flags().String("format", "", "Output format (text, json)")
	hook_approvals_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hook_approvalsCmd.AddCommand(hook_approvals_listCmd)

	carapace.Gen(hook_approvals_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("text", "json"),
	})
}
