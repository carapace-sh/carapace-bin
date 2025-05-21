package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a command within fnm context",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().String("node-dist-mirror", "https://nodejs.org/dist", "<https://nodejs.org/dist/> mirror")
	execCmd.Flags().String("using", "", "Either an explicit version, or a filename with the version written in it")
	execCmd.Flags().String("fnm-dir", "", "The root directory of fnm installation")
	execCmd.Flags().String("log-level", "info", "The log level of fnm commands")
	execCmd.Flags().String("arch", "", "Override the architecture of the installed Node binary. Defaults to arch of fnm binary")
	execCmd.Flags().String("version-file-strategy", "local", "A strategy for how to resolve the Node version. Used whenever `fnm use` or `fnm install` is called without a version, or when `--use-on-cd` is configured on evaluation")
	execCmd.Flags().Bool("corepack-enabled", false, "Enable corepack support for each new installation. This will make fnm call `corepack enable` on every Node.js installation. For more information about corepack see <https://nodejs.org/api/corepack.html>")
	execCmd.Flags().String("resolve-engines", "true", "Resolve `engines.node` field in `package.json` whenever a `.node-version` or `.nvmrc` file is not present. This feature is enabled by default. To disable it, provide `--resolve-engines=false`")
	execCmd.Flags().BoolP("help", "h", false, "Print help (see summary with '-h')")

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"using":                 action.ActionLocalVersions(),
		"log-level":             action.ActionLogLevel(),
		"version-file-strategy": action.ActionVersionFileStrategy(),
		"resolve-engines":       action.ActionResolveEngines(),
	})
}
