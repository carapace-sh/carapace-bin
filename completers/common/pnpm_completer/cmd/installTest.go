package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var installTestCmd = &cobra.Command{
	Use:     "install-test",
	Aliases: []string{"it"},

	Short:   "Runs a `pnpm install` followed immediately by a `pnpm test`",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installTestCmd).Standalone()

	installTestCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	installTestCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects since the specified commit/branch")
	installTestCmd.Flags().String("child-concurrency", "", "Controls the number of child processes run parallelly")
	installTestCmd.Flags().Bool("color", false, "Controls colors in the output")
	installTestCmd.Flags().BoolP("dev", "D", false, "Only `devDependencies` are installed regardless of the `NODE_ENV`")
	installTestCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	installTestCmd.Flags().String("filter", "", "set filter")
	installTestCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	installTestCmd.Flags().Bool("fix-lockfile", false, "Fix broken lockfile entries automatically")
	installTestCmd.Flags().Bool("force", false, "Force reinstall dependencies")
	installTestCmd.Flags().Bool("frozen-lockfile", false, "Don't generate a lockfile and fail if an update is needed")
	installTestCmd.Flags().Bool("global-dir", false, "Specify a custom directory to store global packages")
	installTestCmd.Flags().BoolP("help", "h", false, "Output usage information")
	installTestCmd.Flags().String("hoist-pattern", "", "Hoist all dependencies matching the pattern to")
	installTestCmd.Flags().Bool("ignore-pnpmfile", false, "Disable pnpm hooks defined in .pnpmfile.cjs")
	installTestCmd.Flags().Bool("ignore-scripts", false, "Don't run lifecycle scripts")
	installTestCmd.Flags().String("lockfile-dir", "", "The directory in which the pnpm-lock.yaml will be created")
	installTestCmd.Flags().Bool("lockfile-only", false, "Dependencies are not downloaded")
	installTestCmd.Flags().String("loglevel", "", "What level of logs to report")
	installTestCmd.Flags().String("modules-dir", "", "The directory in which dependencies will be installed")
	installTestCmd.Flags().String("network-concurrency", "", "Maximum number of concurrent network requests")
	installTestCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	installTestCmd.Flags().Bool("no-frozen-lockfile", false, "Don't generate a lockfile and fail if an update is needed")
	installTestCmd.Flags().Bool("no-hoist", false, "Dependencies inside the modules directory will have access only to their listed dependencies")
	installTestCmd.Flags().Bool("no-lockfile", false, "Don't read or generate a `pnpm-lock.yaml` file")
	installTestCmd.Flags().Bool("no-optional", false, "`optionalDependencies` are not installed")
	installTestCmd.Flags().Bool("no-verify-store-integrity", false, "If false, doesn't check whether packages in the store were mutated")
	installTestCmd.Flags().Bool("offline", false, "Trigger an error if any required dependencies are not available in local store")
	installTestCmd.Flags().String("package-import-method", "", "configure import method")
	installTestCmd.Flags().Bool("prefer-frozen-lockfile", false, "If the available `pnpm-lock.yaml` satisfies the `package.json` then perform a headless installation")
	installTestCmd.Flags().Bool("prefer-offline", false, "Skip staleness checks for cached data, but request missing data from the server")
	installTestCmd.Flags().BoolP("prod", "P", false, "Packages in `devDependencies` won't be installed")
	installTestCmd.Flags().String("public-hoist-pattern", "", "Hoist all dependencies matching the pattern to the root of the modules directory")
	installTestCmd.Flags().BoolP("recursive", "r", false, "Run installation recursively in every package found in subdirectories")
	installTestCmd.Flags().String("reporter", "", "configure output")
	installTestCmd.Flags().Bool("shamefully-hoist", false, "All the subdeps will be hoisted into the root node_modules")
	installTestCmd.Flags().Bool("side-effects-cache", false, "Use or cache the results of (pre/post)install hooks")
	installTestCmd.Flags().Bool("side-effects-cache-readonly", false, "Only use the side effects cache if present, do not create it for new packages")
	installTestCmd.Flags().BoolP("silent", "s", false, "No output is logged to the console, except fatal errors")
	installTestCmd.Flags().String("store-dir", "", "The directory in which all the packages are saved on the disk")
	installTestCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	installTestCmd.Flags().Bool("strict-peer-dependencies", false, "Fail on missing or invalid peer dependencies")
	installTestCmd.Flags().String("test-pattern", "", "Defines files related to tests")
	installTestCmd.Flags().Bool("use-running-store-server", false, "Only allows installation with a store server")
	installTestCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	installTestCmd.Flags().Bool("use-store-server", false, "Starts a store server in the background")
	installTestCmd.Flags().Bool("verify-store-integrity", false, "If false, doesn't check whether packages in the store were mutated")
	installTestCmd.Flags().String("virtual-store-dir", "", "The directory with links to the store")
	installTestCmd.Flags().BoolP("workspace-root", "w", false, "Run the command on the root workspace project")
	rootCmd.AddCommand(installTestCmd)

	carapace.Gen(installTestCmd).FlagCompletion(carapace.ActionMap{
		"dir":          carapace.ActionDirectories(),
		"filter":       pnpm.ActionFilter(),
		"filter-prod":  pnpm.ActionFilter(),
		"lockfile-dir": carapace.ActionDirectories(),
		"loglevel":     pnpm.ActionLoglevel(),
		"modules-dir":  carapace.ActionDirectories(),
		"package-import-method": carapace.ActionValuesDescribed(
			"auto", "Clones/hardlinks or copies packages",
			"clone", "Clone (aka copy-on-write) packages from the store",
			"copy", "Copy packages from the store",
			"hardlink", "Hardlink packages from the store",
		),
		"reporter": carapace.ActionValuesDescribed(
			"append-only", "The output is always appended to the end",
			"default", "The default reporter when the stdout is TTY",
			"ndjson", "The most verbose reporter",
			"silent", "No output is logged to the console, except fatal errors",
		),
		"store-dir":         carapace.ActionDirectories(),
		"virtual-store-dir": carapace.ActionDirectories(),
	})
}
