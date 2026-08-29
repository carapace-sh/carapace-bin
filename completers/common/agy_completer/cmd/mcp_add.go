package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
	"github.com/spf13/cobra"
)

var mcpAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add or update an MCP server configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcpAddCmd).Standalone()

	mcpAddCmd.Flags().StringArray("env", []string{}, "Environment variable (KEY=VALUE)")
	mcpAddCmd.Flags().StringArray("header", []string{}, "HTTP header (KEY=VALUE)")
	mcpAddCmd.Flags().String("type", "", "Server type (stdio, http)")

	mcpCmd.AddCommand(mcpAddCmd)

	carapace.Gen(mcpAddCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionExecutables(),
	)

	carapace.Gen(mcpAddCmd).FlagCompletion(carapace.ActionMap{
		"env":    env.ActionNameValues(false),
		"header": http.ActionRequestHeaders(),
		"type":   carapace.ActionValues("stdio", "http"),
	})
}
