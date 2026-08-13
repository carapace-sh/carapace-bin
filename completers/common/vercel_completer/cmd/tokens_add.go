package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tokens_addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create"},
	Short:   "Create a new personal authentication token",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tokens_addCmd).Standalone()

	tokens_addCmd.Flags().String("format", "", "Output format")
	tokens_addCmd.Flags().Bool("json", false, "Output as JSON")
	tokens_addCmd.Flags().String("project", "", "Project name or ID")

	tokensCmd.AddCommand(tokens_addCmd)

	carapace.Gen(tokens_addCmd).FlagCompletion(carapace.ActionMap{
		"format":  carapace.ActionValues("plain", "json"),
		"project": action.ActionProjects(tokens_addCmd),
	})

	carapace.Gen(tokens_addCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json"),
	})
}
