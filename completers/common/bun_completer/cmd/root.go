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
	rootCmd.Flags().String("conditions", "", "Pass custom conditions to resolve")
	rootCmd.Flags().StringP("config", "c", "", "Config file to load bun from")
	rootCmd.Flags().String("console-depth", "", "Set the default depth for console.log object inspection (default: 2)")
	rootCmd.Flags().Bool("cpu-prof", false, "Start CPU profiler and write profile to disk on exit")
	rootCmd.Flags().String("cpu-prof-dir", "", "Specify the directory where the CPU profile will be saved")
	rootCmd.Flags().String("cpu-prof-interval", "", "Specify the sampling interval in microseconds for CPU profiling (default: 1000)")
	rootCmd.Flags().Bool("cpu-prof-md", false, "Output CPU profile in markdown format")
	rootCmd.Flags().String("cpu-prof-name", "", "Specify the name of the CPU profile file")
	rootCmd.Flags().String("cron-period", "", "Cron period for cron execution mode")
	rootCmd.Flags().String("cron-title", "", "Title for cron execution mode")
	rootCmd.Flags().String("cwd", "", "Absolute path to resolve files & entry points from")
	rootCmd.Flags().StringP("define", "d", "", "Substitute K:V while parsing, e.g. --define process.env.NODE_ENV:\"development\"")
	rootCmd.Flags().Bool("disable-hmr", false, "Disable Hot Module Reloading (disables fast refresh too) in bun dev")
	rootCmd.Flags().Bool("disable-react-fast-refresh", false, "Disable React Fast Refresh")
	rootCmd.Flags().String("dns-result-order", "", "Set the default order of DNS lookup results. Valid orders: verbatim (default), ipv4first, ipv6first")
	rootCmd.Flags().String("elide-lines", "", "Number of lines of script output shown when using --filter (default: 10)")
	rootCmd.Flags().String("env-file", "", "Load environment variables from the specified file(s)")
	rootCmd.Flags().Bool("experimental-http2-fetch", false, "Offer h2 in fetch() TLS ALPN")
	rootCmd.Flags().Bool("experimental-http3-fetch", false, "Honor Alt-Svc: h3 in fetch() and upgrade to HTTP/3")
	rootCmd.Flags().Bool("expose-gc", false, "Expose gc() on the global object")
	rootCmd.Flags().String("extension-order", "", "defaults to: .tsx,.ts,.jsx,.js,.json")
	rootCmd.Flags().StringP("external", "e", "", "Exclude module from transpilation (can use * wildcards). ex: -e react")
	rootCmd.Flags().String("fetch-preconnect", "", "Preconnect to a URL while code is loading")
	rootCmd.Flags().Bool("heap-prof", false, "Generate V8 heap snapshot on exit (.heapsnapshot)")
	rootCmd.Flags().String("heap-prof-dir", "", "Specify the directory where the heap profile will be saved")
	rootCmd.Flags().Bool("heap-prof-md", false, "Generate markdown heap profile on exit (for CLI analysis)")
	rootCmd.Flags().String("heap-prof-name", "", "Specify the name of the heap profile file")
	rootCmd.Flags().BoolP("help", "h", false, "Display this help and exit.")
	rootCmd.Flags().Bool("hot", false, "Enable auto reload in bun's JavaScript runtime")
	rootCmd.Flags().BoolS("i", "i", false, "Automatically install dependencies and use global cache in bun's runtime, equivalent to --install=fallback")
	rootCmd.Flags().Bool("if-present", false, "Exit without an error if the entrypoint does not exist")
	rootCmd.Flags().String("import", "", "Alias of --preload, for Node.js compatibility")
	rootCmd.Flags().String("inspect", "", "Activate Bun's debugger")
	rootCmd.Flags().String("inspect-brk", "", "Activate Bun's debugger, set breakpoint on first line of code and wait")
	rootCmd.Flags().String("inspect-wait", "", "Activate Bun's debugger, wait for a connection before executing")
	rootCmd.Flags().String("install", "", "Install dependencies automatically when no node_modules are present")
	rootCmd.Flags().String("jsx-factory", "", "Changes the function called when compiling JSX elements using the classic JSX runtime")
	rootCmd.Flags().String("jsx-fragment", "", "Changes the function called when compiling JSX fragments")
	rootCmd.Flags().String("jsx-import-source", "", "Declares the module specifier to be used for importing the jsx and jsxs factory functions")
	rootCmd.Flags().Bool("jsx-production", false, "Use jsx instead of jsxDEV (default) for the automatic runtime")
	rootCmd.Flags().String("jsx-runtime", "", "\"automatic\" (default) or \"classic\"")
	rootCmd.Flags().StringP("loader", "l", "", "Parse files with .ext:loader, e.g. --loader .js:jsx")
	rootCmd.Flags().String("main-fields", "", "Main fields to lookup in package.json")
	rootCmd.Flags().String("max-http-header-size", "", "Set the maximum size of HTTP headers in bytes. Default is 16KiB")
	rootCmd.Flags().Bool("no-addons", false, "Throw an error if process.dlopen is called, and disable export condition \"node-addons\"")
	rootCmd.Flags().Bool("no-clear-screen", false, "Disable clearing the terminal screen on reload when --hot or --watch is enabled")
	rootCmd.Flags().Bool("no-deprecation", false, "Suppress all reporting of the custom deprecation")
	rootCmd.Flags().Bool("no-env-file", false, "Disable automatic loading of .env files")
	rootCmd.Flags().Bool("no-exit-on-error", false, "Continue running other scripts when one fails (with --parallel/--sequential)")
	rootCmd.Flags().Bool("no-install", false, "Disable auto install in bun's JavaScript runtime")
	rootCmd.Flags().Bool("no-orphans", false, "Exit when the parent process dies, and on exit SIGKILL every descendant")
	rootCmd.Flags().Bool("no-summary", false, "Don't print a summary (when generating .bun")
	rootCmd.Flags().StringP("origin", "u", "", "Rewrite import URLs to start with --origin")
	rootCmd.Flags().Bool("parallel", false, "Run multiple scripts concurrently with Foreman-style output")
	rootCmd.Flags().String("platform", "", "\"bun\" or \"browser\" or \"node\", used when building or bundling")
	rootCmd.Flags().StringP("port", "p", "", "Port to serve bun's dev server on")
	rootCmd.Flags().Bool("prefer-latest", false, "Use the latest matching versions of packages in bun's JavaScript runtime, always checking npm")
	rootCmd.Flags().Bool("prefer-offline", false, "Skip staleness checks for packages in bun's JavaScript runtime and resolve from disk")
	rootCmd.Flags().StringP("preload", "r", "", "Import a module before other modules are loaded")
	rootCmd.Flags().String("public-dir", "", "Top-level directory for .html files, fonts or anything external")
	rootCmd.Flags().Bool("redis-preconnect", false, "Preconnect to $REDIS_URL at startup")
	rootCmd.Flags().String("require", "", "Alias of --preload, for Node.js compatibility")
	rootCmd.Flags().Bool("revision", false, "Print version with revision and exit")
	rootCmd.Flags().Bool("sequential", false, "Run multiple scripts sequentially with Foreman-style output")
	rootCmd.Flags().String("server-bunfile", "", "Use a .server.bun file")
	rootCmd.Flags().String("shell", "", "Control the shell used for package.json scripts. Supports either 'bun' or 'system'")
	rootCmd.Flags().Bool("silent", false, "Don't repeat the command for bun run")
	rootCmd.Flags().Bool("smol", false, "Use less memory, but run garbage collection more often")
	rootCmd.Flags().Bool("sql-preconnect", false, "Preconnect to PostgreSQL at startup")
	rootCmd.Flags().Bool("throw-deprecation", false, "Determine whether or not deprecation warnings result in errors")
	rootCmd.Flags().String("title", "", "Set the process title")
	rootCmd.Flags().String("tsconfig-override", "", "Load tsconfig from path instead of cwd/tsconfig.json")
	rootCmd.Flags().Bool("unhandled-rejections", false, "One of \"strict\", \"throw\", \"warn\", \"none\", or \"warn-with-error-code\"")
	rootCmd.Flags().String("use", "", "Choose a framework, e.g. \"--use next\"")
	rootCmd.Flags().Bool("use-bundled-ca", false, "Use bundled CA store")
	rootCmd.Flags().Bool("use-openssl-ca", false, "Use OpenSSL's default CA store")
	rootCmd.Flags().Bool("use-system-ca", false, "Use the system's trusted certificate authorities")
	rootCmd.Flags().String("user-agent", "", "Set the default User-Agent header for HTTP requests")
	rootCmd.Flags().BoolP("version", "v", false, "Print version and exit")
	rootCmd.Flags().Bool("watch", false, "Automatically restart bun's JavaScript runtime on file change")
	rootCmd.Flags().Bool("workspaces", false, "Run a script in all workspace packages (from the \"workspaces\" field in package.json)")
	rootCmd.Flags().Bool("zero-fill-buffers", false, "Boolean to force Buffer.allocUnsafe(size) to be zero-filled")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bunfile":           carapace.ActionFiles(".bun"),
		"config":            carapace.ActionFiles(".toml"),
		"cpu-prof-dir":      carapace.ActionDirectories(),
		"cwd":               carapace.ActionDirectories(),
		"dns-result-order":  carapace.ActionValues("verbatim", "ipv4first", "ipv6first"),
		"env-file":          carapace.ActionFiles(),
		"extension-order":   carapace.ActionValues(".tsx", ".ts", ".jsx", ".js", ".json").StyleF(style.ForPathExt).UniqueList(","),
		"heap-prof-dir":     carapace.ActionDirectories(),
		"import":            carapace.ActionFiles(),
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
		"preload":        carapace.ActionFiles(),
		"public-dir":     carapace.ActionDirectories(),
		"require":        carapace.ActionFiles(),
		"server-bunfile": carapace.ActionValues(".server.bun"),
		"shell":          carapace.ActionValues("bun", "system"),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(rootCmd.Flag("cwd").Value.String())
	})
}
