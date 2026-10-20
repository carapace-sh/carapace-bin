// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var plugins_installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install plugins from a configured source, Git URL, or archive",
	Aliases: []string{"i", "a", "add"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugins_installCmd).Standalone()

	plugins_installCmd.Flags().BoolP("all", "a", false, "Install every plugin the current config needs that is not installed yet")
	plugins_installCmd.Flags().BoolP("force", "f", false, "Reinstall even if the plugin is already installed")
	plugins_installCmd.Flags().BoolP("help", "h", false, "Print help")
	plugins_installCmd.Flags().StringP("jobs", "j", "", "Number of plugins to install in parallel (default: the `jobs` setting)")
	plugins_installCmd.Flags().CountP("verbose", "v", "Show installation output")
	pluginsCmd.AddCommand(plugins_installCmd)

	carapace.Gen(plugins_installCmd).PositionalCompletion(
		mise.ActionRemotePlugins(),
		http.ActionUrls(),
	)
}
