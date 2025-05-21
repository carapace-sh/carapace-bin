package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a new Node.js version",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().Bool("lts", false, "Install latest LTS")
	installCmd.Flags().String("node-dist-mirror", "https://nodejs.org/dist", "<https://nodejs.org/dist/> mirror")
	installCmd.Flags().String("fnm-dir", "", "The root directory of fnm installation")
	installCmd.Flags().Bool("latest", false, "Install latest version")
	installCmd.Flags().String("progress", "auto", "Show an interactive progress bar for the download status")
	installCmd.Flags().String("log-level", "info", "The log level of fnm commands")
	installCmd.Flags().String("arch", "", "Override the architecture of the installed Node binary. Defaults to arch of fnm binary")
	installCmd.Flags().String("version-file-strategy", "local", "A strategy for how to resolve the Node version. Used whenever `fnm use` or `fnm install` is called without a version, or when `--use-on-cd` is configured on evaluation")
	installCmd.Flags().Bool("corepack-enabled", false, "Enable corepack support for each new installation. This will make fnm call `corepack enable` on every Node.js installation. For more information about corepack see <https://nodejs.org/api/corepack.html>")
	installCmd.Flags().String("resolve-engines", "true", "Resolve `engines.node` field in `package.json` whenever a `.node-version` or `.nvmrc` file is not present. This feature is enabled by default. To disable it, provide `--resolve-engines=false`")
	installCmd.Flags().BoolP("help", "h", false, "Print help (see summary with '-h')")

	carapace.Gen(installCmd).Standalone()

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"progress":              carapace.ActionValues("auto", "never", "always"),
		"log-level":             action.ActionLogLevel(),
		"version-file-strategy": action.ActionVersionFileStrategy(),
		"resolve-engines":       action.ActionResolveEngines(),
	})
}
