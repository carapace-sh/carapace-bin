package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Scaffold a new plugin at ~/.claude/skills/<name>/ (auto-loads next session as <name>@skills-dir)",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_initCmd).Standalone()

	plugin_initCmd.Flags().String("author", "", "Author name")
	plugin_initCmd.Flags().String("author-email", "", "Author email")
	plugin_initCmd.Flags().String("description", "", "Manifest description")
	plugin_initCmd.Flags().BoolP("force", "f", false, "Overwrite an existing .claude-plugin/ at the target")
	plugin_initCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_initCmd.Flags().StringSlice("with", nil, "Also scaffold: skills, agents, hooks, mcp, lsp, output-style, channel")
	pluginCmd.AddCommand(plugin_initCmd)
}
