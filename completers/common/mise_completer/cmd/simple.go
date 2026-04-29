package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func addSimpleCommand(use, short string, aliases ...string) {
	cmd := &cobra.Command{Use: use, Short: short, Aliases: aliases, Run: func(cmd *cobra.Command, args []string) {}}
	carapace.Gen(cmd).Standalone()
	rootCmd.AddCommand(cmd)
}

func init() {
	addSimpleCommand("alias", "Manage aliases")
	addSimpleCommand("backends", "Manage tool backends")
	addSimpleCommand("cache", "Manage mise cache")
	addSimpleCommand("completion [shell]", "Generate shell completions")
	addSimpleCommand("config", "Manage config files")
	addSimpleCommand("current [tool]", "Show current tool versions")
	addSimpleCommand("doctor", "Check mise installation for problems")
	addSimpleCommand("exec [tool@version] -- [command]", "Execute a command with tools installed")
	addSimpleCommand("generate", "Generate mise helper files")
	addSimpleCommand("latest [tool]", "Show latest available version for a tool")
	addSimpleCommand("list [tool]", "List installed versions", "ls")
	addSimpleCommand("ls-remote [tool]", "List remote versions")
	addSimpleCommand("outdated", "Show outdated tools")
	addSimpleCommand("plugins", "Manage plugins")
	addSimpleCommand("prune", "Remove unused versions")
	addSimpleCommand("registry", "List available tools")
	addSimpleCommand("reshim", "Rebuild shims")
	addSimpleCommand("self-update", "Update mise itself")
	addSimpleCommand("settings", "Manage settings")
	addSimpleCommand("sync", "Sync tool versions")
	addSimpleCommand("trust [config]", "Trust a mise config")
	addSimpleCommand("uninstall [tool@version]...", "Uninstall tool versions", "remove", "rm")
	addSimpleCommand("upgrade [tool]...", "Upgrade tools")
	addSimpleCommand("where [tool]", "Show install path for a tool")
	addSimpleCommand("which [bin]", "Show binary path")
}
