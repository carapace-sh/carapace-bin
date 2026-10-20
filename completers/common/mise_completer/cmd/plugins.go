// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pluginsCmd = &cobra.Command{
	Use:     "plugins",
	Short:   "Manage plugins",
	Aliases: []string{"p", "plugin", "plugin-list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginsCmd).Standalone()

	pluginsCmd.Flags().BoolP("all", "a", false, "List all available remote plugins (same as `mise plugins ls-remote`)")
	pluginsCmd.Flags().BoolP("core", "c", false, "List the core tools built into mise instead of installed plugins")
	pluginsCmd.Flags().BoolP("help", "h", false, "Print help")
	pluginsCmd.Flags().Bool("refs", false, "Show each plugin's Git ref and commit, such as `main 1234abc`")
	pluginsCmd.Flags().Bool("url", false, "Show each plugin's Git URL, ref, and commit")
	pluginsCmd.Flags().BoolP("urls", "u", false, "Show each plugin's Git URL, ref, and commit")
	pluginsCmd.Flags().Bool("user", false, "List installed plugins")
	pluginsCmd.Flag("all").Hidden = true
	pluginsCmd.Flag("refs").Hidden = true
	pluginsCmd.Flag("url").Hidden = true
	rootCmd.AddCommand(pluginsCmd)
}
