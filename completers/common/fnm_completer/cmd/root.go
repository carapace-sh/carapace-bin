package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fnm",
	Short: "A fast and simple Node.js manager",
	Long:  "https://github.com/Schniz/fnm",
	Run:   func(*cobra.Command, []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("arch", "", "Override the architecture of the installed Node binary. Defaults to arch of fnm binary")
	rootCmd.PersistentFlags().Bool("corepack-enabled", false, "Enable corepack support for each new installation. This will make fnm call `corepack enable` on every Node.js installation. For more information about corepack see <https://nodejs.org/api/corepack.html>")
	rootCmd.PersistentFlags().String("fnm-dir", "", "The root directory of fnm installation")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print help (see summary with '-h')")
	rootCmd.PersistentFlags().String("log-level", "info", "The log level of fnm commands")
	rootCmd.PersistentFlags().String("node-dist-mirror", "https://nodejs.org/dist", "<https://nodejs.org/dist/> mirror")
	rootCmd.PersistentFlags().String("resolve-engines", "true", "Resolve `engines.node` field in `package.json` whenever a `.node-version` or `.nvmrc` file is not present. This feature is enabled by default. To disable it, provide `--resolve-engines=false`")
	rootCmd.PersistentFlags().String("version-file-strategy", "local", "A strategy for how to resolve the Node version. Used whenever `fnm use` or `fnm install` is called without a version, or when `--use-on-cd` is configured on evaluation")
	rootCmd.Flags().BoolP("version", "v", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"fnm-dir":               carapace.ActionDirectories(),
		"log-level":             action.ActionLogLevel().StyleF(style.ForLogLevel),
		"resolve-engines":       action.ActionResolveEngines(),
		"version-file-strategy": action.ActionVersionFileStrategy(),
	})
}
