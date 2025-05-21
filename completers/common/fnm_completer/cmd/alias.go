package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Alias a version to a common name",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	rootCmd.AddCommand(aliasCmd)

	aliasCmd.Flags().String("node-dist-mirror", "https://nodejs.org/dist", "<https://nodejs.org/dist/> mirror")
	aliasCmd.Flags().String("fnm-dir", "", "The root directory of fnm installation")
	aliasCmd.Flags().String("log-level", "info", "The log level of fnm commands")
	aliasCmd.Flags().String("arch", "", "Override the architecture of the installed Node binary. Defaults to arch of fnm binary")
	aliasCmd.Flags().String("version-file-strategy", "local", "A strategy for how to resolve the Node version. Used whenever `fnm use` or `fnm install` is called without a version, or when `--use-on-cd` is configured on evaluation")
	aliasCmd.Flags().Bool("corepack-enabled", false, "Enable corepack support for each new installation. This will make fnm call `corepack enable` on every Node.js installation. For more information about corepack see <https://nodejs.org/api/corepack.html>")
	aliasCmd.Flags().String("resolve-engines", "true", "Resolve `engines.node` field in `package.json` whenever a `.node-version` or `.nvmrc` file is not present. This feature is enabled by default. To disable it, provide `--resolve-engines=false`")
	aliasCmd.Flags().BoolP("help", "h", false, "Print help (see summary with '-h')")

	carapace.Gen(aliasCmd).Standalone()

	carapace.Gen(aliasCmd).FlagCompletion(carapace.ActionMap{
		"log-level":             action.ActionLogLevel(),
		"version-file-strategy": action.ActionVersionFileStrategy(),
		"resolve-engines":       action.ActionResolveEngines(),
	})

	carapace.Gen(aliasCmd).PositionalCompletion(
		action.ActionLocalVersions(),
	)
}
