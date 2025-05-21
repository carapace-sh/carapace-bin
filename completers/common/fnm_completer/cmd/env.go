package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print and set up required environment variables for fnm",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	rootCmd.AddCommand(envCmd)

	envCmd.Flags().String("node-dist-mirror", "https://nodejs.org/dist", "<https://nodejs.org/dist/> mirror")
	envCmd.Flags().String("shell", "", "The shell syntax to use. Infers when missing")
	envCmd.Flags().String("fnm-dir", "", "The root directory of fnm installation")
	envCmd.Flags().Bool("json", false, "Print JSON instead of shell commands")
	envCmd.Flags().String("log-level", "info", "The log level of fnm commands")
	envCmd.Flags().Bool("use-on-cd", false, "Print the script to change Node versions every directory change")
	envCmd.Flags().String("arch", "", "Override the architecture of the installed Node binary. Defaults to arch of fnm binary")
	envCmd.Flags().String("version-file-strategy", "local", "A strategy for how to resolve the Node version. Used whenever `fnm use` or `fnm install` is called without a version, or when `--use-on-cd` is configured on evaluation")
	envCmd.Flags().Bool("corepack-enabled", false, "Enable corepack support for each new installation. This will make fnm call `corepack enable` on every Node.js installation. For more information about corepack see <https://nodejs.org/api/corepack.html>")
	envCmd.Flags().String("resolve-engines", "true", "Resolve `engines.node` field in `package.json` whenever a `.node-version` or `.nvmrc` file is not present. This feature is enabled by default. To disable it, provide `--resolve-engines=false`")
	envCmd.Flags().BoolP("help", "h", false, "Print help (see summary with '-h')")

	carapace.Gen(envCmd).Standalone()

	carapace.Gen(envCmd).FlagCompletion(carapace.ActionMap{
		"shell":                 carapace.ActionValues("bash", "zsh", "fish", "powershell"),
		"log-level":             action.ActionLogLevel(),
		"version-file-strategy": action.ActionVersionFileStrategy(),
		"resolve-engines":       action.ActionResolveEngines(),
	})
}
