package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install the tools in your config, or specific tool versions",
	Aliases: []string{"i"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().String("before", "", "Only install versions released before a date or at least a duration ago")
	installCmd.Flags().BoolP("dry-run", "n", false, "Show what would be installed, without installing anything")
	installCmd.Flags().Bool("dry-run-code", false, "Like --dry-run, but exit with status 1 when anything needs installing")
	installCmd.Flags().BoolP("force", "f", false, "Reinstall even if already installed")
	installCmd.Flags().BoolP("help", "h", false, "Print help")
	installCmd.Flags().Bool("include-lazy", false, "Also install tools configured with `lazy = true`")
	installCmd.Flags().Bool("include-task-tools", false, "Also install tools required by tasks in the current scope")
	installCmd.Flags().StringP("jobs", "j", "", "How many tools to install in parallel; defaults to the `jobs` setting")
	installCmd.Flags().String("minimum-release-age", "", "Only install versions released before a date or at least a duration ago")
	installCmd.Flags().Bool("monorepo", false, "Also install tools from every config root in `[monorepo].config_roots`")
	installCmd.Flags().Bool("raw", false, "Connect install commands directly to the terminal; implies --jobs=1")
	installCmd.Flags().String("shared", "", "Install into DIR instead of the mise data directory")
	installCmd.Flags().Bool("system", false, "Install into the system-wide shared directory")
	installCmd.Flags().CountP("verbose", "v", "Show installer output")
	installCmd.Flag("before").Hidden = true
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"shared": carapace.ActionDirectories(),
	})

	carapace.Gen(installCmd).PositionalAnyCompletion(
		mise.ActionToolVersions(),
	)
}
