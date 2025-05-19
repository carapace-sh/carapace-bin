package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i"},
	Short:   "Installs all dependencies of the project in the current working directory",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	installCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects since the specified commit/branch")
	installCmd.Flags().String("child-concurrency", "", "Controls the number of child processes run parallelly")
	installCmd.Flags().Bool("color", false, "Controls colors in the output")
	installCmd.Flags().BoolP("dev", "D", false, "Only `devDependencies` are installed regardless of the `NODE_ENV`")
	installCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	installCmd.Flags().String("filter", "", "set filter")
	installCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	installCmd.Flags().Bool("fix-lockfile", false, "Fix broken lockfile entries automatically")
	installCmd.Flags().Bool("force", false, "Force reinstall dependencies")
	installCmd.Flags().Bool("frozen-lockfile", false, "Don't generate a lockfile and fail if an update is needed")
	installCmd.Flags().Bool("global-dir", false, "Specify a custom directory to store global packages")
	installCmd.Flags().BoolP("help", "h", false, "Output usage information")
	installCmd.Flags().String("hoist-pattern", "", "Hoist all dependencies matching the pattern to")
	installCmd.Flags().Bool("ignore-pnpmfile", false, "Disable pnpm hooks defined in .pnpmfile.cjs")
	installCmd.Flags().Bool("ignore-scripts", false, "Don't run lifecycle scripts")
	installCmd.Flags().String("lockfile-dir", "", "The directory in which the pnpm-lock.yaml will be created")
	installCmd.Flags().Bool("lockfile-only", false, "Dependencies are not downloaded")
	installCmd.Flags().String("loglevel", "", "What level of logs to report")
	installCmd.Flags().String("modules-dir", "", "The directory in which dependencies will be installed")
	installCmd.Flags().String("network-concurrency", "", "Maximum number of concurrent network requests")
	installCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	installCmd.Flags().Bool("no-frozen-lockfile", false, "Don't generate a lockfile and fail if an update is needed")
	installCmd.Flags().Bool("no-hoist", false, "Dependencies inside the modules directory will have access only to their listed dependencies")
	installCmd.Flags().Bool("no-lockfile", false, "Don't read or generate a `pnpm-lock.yaml` file")
	installCmd.Flags().Bool("no-optional", false, "`optionalDependencies` are not installed")
	installCmd.Flags().Bool("no-verify-store-integrity", false, "If false, doesn't check whether packages in the store were mutated")
	installCmd.Flags().Bool("offline", false, "Trigger an error if any required dependencies are not available in local store")
	installCmd.Flags().String("package-import-method", "", "configure import method")
	installCmd.Flags().Bool("prefer-frozen-lockfile", false, "If the available `pnpm-lock.yaml` satisfies the `package.json` then perform a headless installation")
	installCmd.Flags().Bool("prefer-offline", false, "Skip staleness checks for cached data, but request missing data from the server")
	installCmd.Flags().BoolP("prod", "P", false, "Packages in `devDependencies` won't be installed")
	installCmd.Flags().String("public-hoist-pattern", "", "Hoist all dependencies matching the pattern to the root of the modules directory")
	installCmd.Flags().BoolP("recursive", "r", false, "Run installation recursively in every package found in subdirectories")
	installCmd.Flags().String("reporter", "", "configure output")
	installCmd.Flags().Bool("shamefully-hoist", false, "All the subdeps will be hoisted into the root node_modules")
	installCmd.Flags().Bool("side-effects-cache", false, "Use or cache the results of (pre/post)install hooks")
	installCmd.Flags().Bool("side-effects-cache-readonly", false, "Only use the side effects cache if present, do not create it for new packages")
	installCmd.Flags().BoolP("silent", "s", false, "No output is logged to the console, except fatal errors")
	installCmd.Flags().String("store-dir", "", "The directory in which all the packages are saved on the disk")
	installCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	installCmd.Flags().Bool("strict-peer-dependencies", false, "Fail on missing or invalid peer dependencies")
	installCmd.Flags().String("test-pattern", "", "Defines files related to tests")
	installCmd.Flags().Bool("use-running-store-server", false, "Only allows installation with a store server")
	installCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	installCmd.Flags().Bool("use-store-server", false, "Starts a store server in the background")
	installCmd.Flags().Bool("verify-store-integrity", false, "If false, doesn't check whether packages in the store were mutated")
	installCmd.Flags().String("virtual-store-dir", "", "The directory with links to the store")
	installCmd.Flags().BoolP("workspace-root", "w", false, "Run the command on the root workspace project")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
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
