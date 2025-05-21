package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Change Node.js version",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(useCmd).Standalone()

	useCmd.Flags().Bool("install-if-missing", false, "Install the version if it isn't installed yet")
	useCmd.Flags().String("node-dist-mirror", "https://nodejs.org/dist", "<https://nodejs.org/dist/> mirror")
	useCmd.Flags().String("fnm-dir", "", "The root directory of fnm installation")
	useCmd.Flags().Bool("silent-if-unchanged", false, "Don't output a message identifying the version being used if it will not change due to execution of this command")
	useCmd.Flags().String("log-level", "info", "The log level of fnm commands")
	useCmd.Flags().String("arch", "", "Override the architecture of the installed Node binary. Defaults to arch of fnm binary")
	useCmd.Flags().String("version-file-strategy", "local", "A strategy for how to resolve the Node version. Used whenever `fnm use` or `fnm install` is called without a version, or when `--use-on-cd` is configured on evaluation")
	useCmd.Flags().Bool("corepack-enabled", false, "Enable corepack support for each new installation. This will make fnm call `corepack enable` on every Node.js installation. For more information about corepack see <https://nodejs.org/api/corepack.html>")
	useCmd.Flags().String("resolve-engines", "true", "Resolve `engines.node` field in `package.json` whenever a `.node-version` or `.nvmrc` file is not present. This feature is enabled by default. To disable it, provide `--resolve-engines=false`")
	useCmd.Flags().BoolP("help", "h", false, "Print help (see summary with '-h')")

	carapace.Gen(useCmd).FlagCompletion(carapace.ActionMap{
		"log-level":             action.ActionLogLevel(),
		"version-file-strategy": action.ActionVersionFileStrategy(),
		"resolve-engines":       action.ActionResolveEngines(),
	})

	carapace.Gen(useCmd).PositionalAnyCompletion(
		action.ActionInstalledVersions(),
	)
}
