package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugins_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List installed plugins",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugins_lsCmd).Standalone()

	plugins_lsCmd.Flags().BoolP("all", "a", false, "List all available remote plugins (same as `mise plugins ls-remote`)")
	plugins_lsCmd.Flags().BoolP("core", "c", false, "List the core tools built into mise instead of installed plugins")
	plugins_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	plugins_lsCmd.Flags().BoolP("outdated", "o", false, "Only show plugins whose Git remote has newer commits")
	plugins_lsCmd.Flags().Bool("refs", false, "Show each plugin's Git ref and commit, such as `main 1234abc`")
	plugins_lsCmd.Flags().Bool("url", false, "Show each plugin's Git URL, ref, and commit")
	plugins_lsCmd.Flags().BoolP("urls", "u", false, "Show each plugin's Git URL, ref, and commit")
	plugins_lsCmd.Flags().Bool("user", false, "List installed plugins")
	plugins_lsCmd.Flag("all").Hidden = true
	plugins_lsCmd.Flag("refs").Hidden = true
	plugins_lsCmd.Flag("url").Hidden = true
	pluginsCmd.AddCommand(plugins_lsCmd)
}
