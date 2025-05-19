package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "bun",
	Short: "a fast bundler, transpiler, JavaScript Runtime and package manager for web software",
	Long:  "https://bun.sh/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("bun", "b", false, "Force a script or package to use Bun.js instead of Node.js")
	rootCmd.Flags().String("bunfile", "", "Use a .bun file")
	rootCmd.Flags().StringP("config", "c", "", "Config file to load bun from")
	rootCmd.Flags().String("cwd", "", "Absolute path to resolve files & entry points from")
	rootCmd.Flags().StringP("define", "d", "", "Substitute K:V while parsing, e.g. --define process.env.NODE_ENV:\"development\"")
	rootCmd.Flags().Bool("disable-hmr", false, "Disable Hot Module Reloading (disables fast refresh too) in bun dev")
	rootCmd.Flags().Bool("disable-react-fast-refresh", false, "Disable React Fast Refresh")
	rootCmd.Flags().String("extension-order", "", "defaults to: .tsx,.ts,.jsx,.js,.json")
	rootCmd.Flags().StringP("external", "e", "", "Exclude module from transpilation (can use * wildcards). ex: -e react")
	rootCmd.Flags().BoolP("help", "h", false, "Display this help and exit.")
	rootCmd.Flags().Bool("hot", false, "Enable auto reload in bun's JavaScript runtime")
	rootCmd.Flags().BoolS("i", "i", false, "Automatically install dependencies and use global cache in bun's runtime, equivalent to --install=fallback")
	rootCmd.Flags().String("install", "", "Install dependencies automatically when no node_modules are present")
	rootCmd.Flags().String("jsx-factory", "", "Changes the function called when compiling JSX elements using the classic JSX runtime")
	rootCmd.Flags().String("jsx-fragment", "", "Changes the function called when compiling JSX fragments")
	rootCmd.Flags().String("jsx-import-source", "", "Declares the module specifier to be used for importing the jsx and jsxs factory functions")
	rootCmd.Flags().Bool("jsx-production", false, "Use jsx instead of jsxDEV (default) for the automatic runtime")
	rootCmd.Flags().String("jsx-runtime", "", "\"automatic\" (default) or \"classic\"")
	rootCmd.Flags().StringP("loader", "l", "", "Parse files with .ext:loader, e.g. --loader .js:jsx")
	rootCmd.Flags().String("main-fields", "", "Main fields to lookup in package.json")
	rootCmd.Flags().Bool("no-install", false, "Disable auto install in bun's JavaScript runtime")
	rootCmd.Flags().Bool("no-summary", false, "Don't print a summary (when generating .bun")
	rootCmd.Flags().StringP("origin", "u", "", "Rewrite import URLs to start with --origin")
	rootCmd.Flags().String("platform", "", "\"bun\" or \"browser\" or \"node\", used when building or bundling")
	rootCmd.Flags().StringP("port", "p", "", "Port to serve bun's dev server on")
	rootCmd.Flags().Bool("prefer-latest", false, "Use the latest matching versions of packages in bun's JavaScript runtime, always checking npm")
	rootCmd.Flags().Bool("prefer-offline", false, "Skip staleness checks for packages in bun's JavaScript runtime and resolve from disk")
	rootCmd.Flags().StringP("preload", "r", "", "Import a module before other modules are loaded")
	rootCmd.Flags().String("public-dir", "", "Top-level directory for .html files, fonts or anything external")
	rootCmd.Flags().String("server-bunfile", "", "Use a .server.bun file")
	rootCmd.Flags().Bool("silent", false, "Don't repeat the command for bun run")
	rootCmd.Flags().String("tsconfig-override", "", "Load tsconfig from path instead of cwd/tsconfig.json")
	rootCmd.Flags().String("use", "", "Choose a framework, e.g. \"--use next\"")
	rootCmd.Flags().BoolP("version", "v", false, "Print version and exit")
	rootCmd.Flags().Bool("watch", false, "Automatically restart bun's JavaScript runtime on file change")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bunfile":           carapace.ActionFiles(".bun"),
		"config":            carapace.ActionFiles(".toml"),
		"cwd":               carapace.ActionDirectories(),
		"extension-order":   carapace.ActionValues(".tsx", ".ts", ".jsx", ".js", ".json").StyleF(style.ForPathExt).UniqueList(","),
		"install":           carapace.ActionValues("auto", "force", "fallback").StyleF(style.ForKeyword),
		"jsx-import-source": carapace.ActionValues("react"),
		"jsx-runtime":       carapace.ActionValues("automatic", "classic"),
		"loader": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return fs.ActionFilenameExtensions().Suffix(":")
			case 1:
				return carapace.ActionValues("jsx", "js", "json", "tsx", "ts", "css").StyleF(func(s string, sc style.Context) string {
					return style.ForPathExt("."+s, sc)
				})
			default:
				return carapace.ActionValues()
			}
		}).UniqueList(","),
		"platform":       carapace.ActionValues("bun", "browser", "node"),
		"port":           net.ActionPorts(),
		"public-dir":     carapace.ActionDirectories(),
		"server-bunfile": carapace.ActionValues(".server.bun"),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(rootCmd.Flag("cwd").Value.String())
	})
}
