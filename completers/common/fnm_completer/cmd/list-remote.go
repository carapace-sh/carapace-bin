package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listRemoteCmd = &cobra.Command{
	Use:     "list-remote",
	Short:   "List all remote Node.js versions",
	Aliases: []string{"ls-remote"},
	Run:     func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(listRemoteCmd).Standalone()

	listRemoteCmd.Flags().String("filter", "", "Filter versions by a user-defined version or a semver range")
	listRemoteCmd.Flags().String("node-dist-mirror", "https://nodejs.org/dist", "<https://nodejs.org/dist/> mirror")
	listRemoteCmd.Flags().String("fnm-dir", "", "The root directory of fnm installation")
	listRemoteCmd.Flags().String("lts", "", "Show only LTS versions (optionally filter by LTS codename)")
	listRemoteCmd.Flags().String("sort", "asc", "Version sorting order")
	listRemoteCmd.Flags().Bool("latest", false, "Only show the latest matching version")
	listRemoteCmd.Flags().String("log-level", "info", "The log level of fnm commands")
	listRemoteCmd.Flags().String("arch", "", "Override the architecture of the installed Node binary. Defaults to arch of fnm binary")
	listRemoteCmd.Flags().String("version-file-strategy", "local", "A strategy for how to resolve the Node version. Used whenever `fnm use` or `fnm install` is called without a version, or when `--use-on-cd` is configured on evaluation")
	listRemoteCmd.Flags().Bool("corepack-enabled", false, "Enable corepack support for each new installation. This will make fnm call `corepack enable` on every Node.js installation. For more information about corepack see <https://nodejs.org/api/corepack.html>")
	listRemoteCmd.Flags().String("resolve-engines", "true", "Resolve `engines.node` field in `package.json` whenever a `.node-version` or `.nvmrc` file is not present. This feature is enabled by default. To disable it, provide `--resolve-engines=false`")
	listRemoteCmd.Flags().BoolP("help", "h", false, "Print help (see summary with '-h')")

	carapace.Gen(listRemoteCmd).FlagCompletion(carapace.ActionMap{
		"sort":                  carapace.ActionValues("asc", "desc"),
		"log-level":             action.ActionLogLevel(),
		"version-file-strategy": action.ActionVersionFileStrategy(),
		"resolve-engines":       action.ActionResolveEngines(),
	})
}
