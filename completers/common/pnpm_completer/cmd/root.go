package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "pnpm",
	Short: "Experimental package manager for node.js",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().Bool("aggregate-output", false, "Hold each script's streamed output until the script exits, then print it as one block")
	rootCmd.PersistentFlags().Bool("bail", false, "Stop a recursive command after the first failure")
	rootCmd.PersistentFlags().StringSlice("changed-files-ignore-pattern", nil, "Glob patterns of changed files that the `[since]` `--filter` selector should ignore")
	rootCmd.PersistentFlags().String("color", "", "Force colored output")
	rootCmd.PersistentFlags().StringP("dir", "C", ".", "Set working directory. Accepted anywhere on the command line, before or after the subcommand, like every other rc-option")
	rootCmd.PersistentFlags().Bool("fail-if-no-match", false, "Exit with code 1 when the `--filter` / `--filter-prod` selectors match no workspace project")
	rootCmd.PersistentFlags().StringSliceP("filter", "F", nil, "Select which workspace projects to run on. Repeat to add more. Each selector can be a name pattern (`@scope/*`), a path (`./pkg`), a dependency query (`foo...`), an exclusion (`!bar`), a directory (`{dir}`), or a changed-since query (`[since]`)")
	rootCmd.PersistentFlags().StringSlice("filter-prod", nil, "Like `--filter`, but follow only production dependencies when selecting projects")
	rootCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.PersistentFlags().String("http-proxy", "", "Proxy for HTTP registry and tarball requests")
	rootCmd.PersistentFlags().String("https-proxy", "", "Proxy for HTTPS registry and tarball requests")
	rootCmd.Flags().Bool("if-present", false, "Don't fail when the named script is undefined")
	rootCmd.PersistentFlags().Bool("ignore-workspace", false, "Run as if the project were standalone, ignoring any `pnpm-workspace.yaml` above it")
	rootCmd.PersistentFlags().Bool("include-workspace-root", false, "Also run a recursive command on the root workspace project, which `run` / `exec` / `add` / `test` otherwise leave out")
	rootCmd.PersistentFlags().String("loglevel", "", "What level of logs to print. Mirrors pnpm's universal `--loglevel` option: `silent` selects the silent reporter over any `--reporter` choice; the other levels cap the default reporter's output")
	rootCmd.PersistentFlags().Bool("no-bail", false, "Recursive only: keep going after a project fails")
	rootCmd.PersistentFlags().Bool("no-color", false, "Disable colored output")
	rootCmd.PersistentFlags().Bool("no-include-workspace-root", false, "Leave the root workspace project out of a recursive command, overriding an `includeWorkspaceRoot: true` setting")
	rootCmd.PersistentFlags().String("no-proxy", "", "Hosts that bypass configured proxies")
	rootCmd.PersistentFlags().Bool("no-reporter-hide-prefix", false, "Prefix the streamed output of running scripts with the project it came from, overriding a `reporterHidePrefix: true` setting")
	rootCmd.PersistentFlags().Bool("no-reverse", false, "Process recursive workspace projects in their normal order")
	rootCmd.PersistentFlags().Bool("no-sort", false, "Run recursive workspace projects in workspace order")
	rootCmd.PersistentFlags().String("npmrc-auth-file", "", "Path to an `.npmrc` to read auth settings from, overriding the default `~/.npmrc`")
	rootCmd.PersistentFlags().Bool("parallel", false, "Run scripts in every selected workspace project concurrently, disregarding topological sorting")
	rootCmd.PersistentFlags().String("prefix", ".", "Set working directory. Accepted anywhere on the command line, before or after the subcommand, like every other rc-option")
	rootCmd.PersistentFlags().BoolP("recursive", "r", false, "Run the command for every project in the workspace instead of only the project in `--dir`")
	rootCmd.PersistentFlags().String("registry", "", "Base URL of the npm registry to resolve and fetch packages from. Universal rc-option: accepted on every command and layered onto the config like `--config.registry=<url>`. Commands that expose their own `--registry` still read the same value")
	rootCmd.PersistentFlags().Bool("report-summary", false, "Recursive only: write a `pnpm-exec-summary.json` execution report")
	rootCmd.PersistentFlags().String("reporter", "default", "Reporter output format")
	rootCmd.PersistentFlags().Bool("reporter-hide-prefix", false, "Omit the project prefix from the streamed output of running scripts. A `run` / `exec` option pnpm accepts anywhere on the command line, like the recursive-run flags above")
	rootCmd.PersistentFlags().String("resume-from", "", "Recursive only: resume execution from the given package")
	rootCmd.PersistentFlags().Bool("reverse", false, "Process recursive workspace projects in reverse order")
	rootCmd.PersistentFlags().Bool("sort", false, "Keep recursive workspace projects sorted topologically")
	rootCmd.PersistentFlags().String("state-dir", "", "Directory in which pnpm persists machine-local state")
	rootCmd.PersistentFlags().String("store", "", "Directory in which the package store is created. Relative paths are resolved from the workspace root, or from `--dir` outside a workspace")
	rootCmd.PersistentFlags().String("store-dir", "", "Directory in which the package store is created. Relative paths are resolved from the workspace root, or from `--dir` outside a workspace")
	rootCmd.PersistentFlags().Bool("stream", false, "Stream a recursive command's script output as it arrives, one prefixed line at a time")
	rootCmd.PersistentFlags().StringSlice("test-pattern", nil, "Glob patterns naming test files, used by the `[since]` `--filter` selector to decide which changes count")
	rootCmd.PersistentFlags().Bool("use-stderr", false, "Divert the reporter's output to stderr, leaving stdout for the command's own result")
	rootCmd.PersistentFlags().String("userconfig", "", "Path to an `.npmrc` to read auth settings from, overriding the default `~/.npmrc`")
	rootCmd.Flags().BoolP("version", "v", false, "Print the pnpm version")
	rootCmd.PersistentFlags().String("workspace-concurrency", "", "Maximum number of workspace projects to process in parallel")
	rootCmd.PersistentFlags().StringSlice("workspace-packages", nil, "Glob patterns selecting the workspace's projects, overriding the `packages` field of `pnpm-workspace.yaml`. Repeat to add more")
	rootCmd.PersistentFlags().BoolP("workspace-root", "w", false, "Run the command on the root workspace project")
	rootCmd.PersistentFlags().BoolP("yes", "y", false, "Automatically answer yes to prompts")
	rootCmd.Flag("bail").Hidden = true
	rootCmd.Flag("color").NoOptDefVal = " "
	rootCmd.Flag("if-present").Hidden = true
	rootCmd.Flag("no-bail").Hidden = true
	rootCmd.Flag("no-color").Hidden = true
	rootCmd.Flag("no-reporter-hide-prefix").Hidden = true
	rootCmd.Flag("no-reverse").Hidden = true
	rootCmd.Flag("prefix").Hidden = true
	rootCmd.Flag("report-summary").Hidden = true
	rootCmd.Flag("reporter-hide-prefix").Hidden = true
	rootCmd.Flag("resume-from").Hidden = true
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dir":             carapace.ActionDirectories(),
		"filter":          pnpm.ActionFilters(),
		"filter-prod":     pnpm.ActionFilters(),
		"loglevel":        pnpm.ActionLoglevels(),
		"npmrc-auth-file": carapace.ActionFiles(),
		"prefix":          carapace.ActionDirectories(),
		"reporter":        carapace.ActionValuesDescribed("default", "Rich visual output: a progress line, a packages diff, lifecycle output, and a `Done in ...` summary", "append-only", "Like `default` but forces the append-only rendering even on a TTY", "ndjson", "Newline-delimited JSON on stderr", "silent", "No progress output"),
		"state-dir":       carapace.ActionDirectories(),
		"store":           carapace.ActionDirectories(),
		"store-dir":       carapace.ActionDirectories(),
		"userconfig":      carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.ChdirF(traverse.Flag(cmd.Flag("dir")))
	})
}
