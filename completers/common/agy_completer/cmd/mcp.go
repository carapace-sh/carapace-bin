package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"strings"
)

var mcpCmd = &cobra.Command{
	Use:     "mcp",
	GroupID: "integration",
	Short:   "Manage MCP servers (add, remove, list, enable, disable)",
	Run:     func(cmd *cobra.Command, args []string) {},
}

var mcpAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add or update an MCP server configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var mcpRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an MCP server configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var mcpListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured MCP servers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var mcpEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable an MCP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var mcpDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable an MCP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mcpCmd).Standalone()
	carapace.Gen(mcpAddCmd).Standalone()
	carapace.Gen(mcpRemoveCmd).Standalone()
	carapace.Gen(mcpListCmd).Standalone()
	carapace.Gen(mcpEnableCmd).Standalone()
	carapace.Gen(mcpDisableCmd).Standalone()

	mcpCmd.AddCommand(
		mcpAddCmd,
		mcpRemoveCmd,
		mcpListCmd,
		mcpEnableCmd,
		mcpDisableCmd,
	)

	rootCmd.AddCommand(mcpCmd)

	carapace.Gen(mcpAddCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionExecutables(),
	)

	carapace.Gen(mcpDisableCmd).PositionalCompletion(
		ActionMcpServers(),
	)

	carapace.Gen(mcpEnableCmd).PositionalCompletion(
		ActionMcpServers(),
	)

	carapace.Gen(mcpRemoveCmd).PositionalCompletion(
		ActionMcpServers(),
	)
}

func ActionMcpServers() carapace.Action {
	return carapace.ActionExecCommand("agy", "mcp", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if line == "" || strings.Contains(line, "No MCP servers") || strings.HasPrefix(line, "NAME") {
				continue
			}

			parts := strings.Fields(line)
			if len(parts) > 0 {
				vals = append(vals, parts[0])
			}
		}
		return carapace.ActionValues(vals...)
	})
}
