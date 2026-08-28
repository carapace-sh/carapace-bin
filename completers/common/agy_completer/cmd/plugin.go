package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"strings"
)

var pluginCmd = &cobra.Command{
	Use:     "plugin",
	GroupID: "integration",
	Short:   "Manage plugins (install, uninstall, list, enable, disable)",
	Run:     func(cmd *cobra.Command, args []string) {},
	Aliases: []string{"plugins"},
}

var pluginListCmd = &cobra.Command{
	Use:   "list",
	Short: "List imported plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pluginImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import plugins from gemini or claude",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pluginInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a plugin (supports plugin@marketplace)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pluginUninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pluginEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pluginDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pluginValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pluginLinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Generate link to a marketplace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pluginHelpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show the plugin command list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginCmd).Standalone()
	carapace.Gen(pluginListCmd).Standalone()
	carapace.Gen(pluginImportCmd).Standalone()
	carapace.Gen(pluginInstallCmd).Standalone()
	carapace.Gen(pluginUninstallCmd).Standalone()
	carapace.Gen(pluginEnableCmd).Standalone()
	carapace.Gen(pluginDisableCmd).Standalone()
	carapace.Gen(pluginValidateCmd).Standalone()
	carapace.Gen(pluginLinkCmd).Standalone()
	carapace.Gen(pluginHelpCmd).Standalone()

	pluginCmd.AddCommand(
		pluginListCmd,
		pluginImportCmd,
		pluginInstallCmd,
		pluginUninstallCmd,
		pluginEnableCmd,
		pluginDisableCmd,
		pluginValidateCmd,
		pluginLinkCmd,
		pluginHelpCmd,
	)

	rootCmd.AddCommand(pluginCmd)

	carapace.Gen(pluginDisableCmd).PositionalCompletion(
		ActionPlugins(),
	)

	carapace.Gen(pluginEnableCmd).PositionalCompletion(
		ActionPlugins(),
	)

	carapace.Gen(pluginImportCmd).PositionalCompletion(
		carapace.ActionValues("gemini", "claude"),
	)

	carapace.Gen(pluginInstallCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(pluginLinkCmd).PositionalCompletion(
		// TODO: update positional completion once documentation for agy plugin link is available
		carapace.ActionValues().Usage("marketplace name"),
		carapace.ActionValues().Usage("target"),
	)

	carapace.Gen(pluginUninstallCmd).PositionalCompletion(
		ActionPlugins(),
	)

	carapace.Gen(pluginValidateCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}

func ActionPlugins() carapace.Action {
	return carapace.ActionExecCommand("agy", "plugin", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if line == "" || strings.Contains(line, "No imported plugins") || strings.HasPrefix(line, "NAME") {
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
