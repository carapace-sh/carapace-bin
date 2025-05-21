package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a Node.js version",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().String("node-dist-mirror", "https://nodejs.org/dist", "<https://nodejs.org/dist/> mirror")
	uninstallCmd.Flags().String("using", "", "Either an explicit version, or a filename with the version written in it")
	uninstallCmd.Flags().String("fnm-dir", "", "The root directory of fnm installation")
	uninstallCmd.Flags().String("log-level", "info", "The log level of fnm commands")
	uninstallCmd.Flags().String("arch", "", "Override the architecture of the installed Node binary. Defaults to arch of fnm binary")
	uninstallCmd.Flags().String("version-file-strategy", "local", "A strategy for how to resolve the Node version. Used whenever `fnm use` or `fnm install` is called without a version, or when `--use-on-cd` is configured on evaluation")
	uninstallCmd.Flags().Bool("corepack-enabled", false, "Enable corepack support for each new installation. This will make fnm call `corepack enable` on every Node.js installation. For more information about corepack see <https://nodejs.org/api/corepack.html>")
	uninstallCmd.Flags().String("resolve-engines", "true", "Resolve `engines.node` field in `package.json` whenever a `.node-version` or `.nvmrc` file is not present. This feature is enabled by default. To disable it, provide `--resolve-engines=false`")
	uninstallCmd.Flags().BoolP("help", "h", false, "Print help (see summary with '-h')")

	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"using":                 action.ActionLocalVersions(),
		"log-level":             action.ActionLogLevel(),
		"version-file-strategy": action.ActionVersionFileStrategy(),
		"resolve-engines":       action.ActionResolveEngines(),
	})

	// TESTME:
	carapace.Gen(unaliasCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionLocalVersions().Invoke(c).Merge(
				action.ActionAliases().Invoke(c),
			).ToA()
		}),
	)
}
