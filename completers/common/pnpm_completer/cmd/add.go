package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().StringSlice("allow-build", nil, "Package names allowed to run lifecycle (build) scripts during this install, appended to `allowBuilds`. Prefix a name with `!` to deny its scripts instead. May be repeated")
	addCmd.Flags().Bool("auto-dedupe", false, "Deduplicate compatible dependency versions during installation")
	addCmd.Flags().Bool("config", false, "Add the package as a configuration dependency")
	addCmd.Flags().StringSlice("cpu", nil, "CPU architectures whose platform-specific optional dependencies should be installed. Repeat or comma-separate for multiple values")
	addCmd.Flags().Bool("dev", false, "Leave dependencies and optionalDependencies out of `node_modules`")
	addCmd.Flags().Bool("force", false, "Reinstall every package the lockfile names: relink packages an earlier install already materialized, and install optional dependencies whose `cpu` / `os` / `libc` / `engines` don't match the host instead of skipping them")
	addCmd.Flags().BoolP("global", "g", false, "Install the package globally, linking its bins into the global bin directory")
	addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	addCmd.Flags().Bool("ignore-pnpmfile", false, "Disable pnpm hooks defined in `.pnpmfile.cjs`, including the pnpmfiles of config dependencies")
	addCmd.Flags().Bool("ignore-scripts", false, "Don't run lifecycle scripts of the project or its dependencies. Packages are still installed; only their build scripts are skipped, and the install won't fail because of it")
	addCmd.Flags().Bool("ignore-workspace-root-check", false, "Permit adding dependencies to a multi-package workspace root without `-w`")
	addCmd.Flags().StringSlice("libc", nil, "libc families whose platform-specific optional dependencies should be installed (`glibc`, `musl`). Repeat or comma-separate for multiple values")
	addCmd.Flags().String("lockfile-dir", "", "The directory in which `pnpm-lock.yaml` is created. Several projects may share a single lockfile")
	addCmd.Flags().Bool("lockfile-only", false, "Dependencies are not downloaded. Only `pnpm-lock.yaml` is updated")
	addCmd.Flags().Bool("no-auto-dedupe", false, "Disable automatic deduplication configured in pnpm-workspace.yaml")
	addCmd.Flags().Bool("no-ignore-scripts", false, "Run lifecycle scripts even when the configuration disables them")
	addCmd.Flags().Bool("no-ignore-workspace-root-check", false, "Keep the workspace-root safety check enabled")
	addCmd.Flags().Bool("no-optional", false, "Exclude optionalDependencies while materializing the updated project")
	addCmd.Flags().Bool("no-save-peer", false, "Don't add the packages to peerDependencies, overriding a `savePeer: true` setting")
	addCmd.Flags().Bool("no-save-types", false, "Do not add `@types` packages automatically")
	addCmd.Flags().Bool("optional", false, "Include optionalDependencies while materializing the updated project")
	addCmd.Flags().StringSlice("os", nil, "Operating systems whose platform-specific optional dependencies should be installed. Repeat or comma-separate for multiple values")
	addCmd.Flags().Bool("prod", false, "Leave devDependencies out of `node_modules`")
	addCmd.Flags().Bool("production", false, "Leave devDependencies out of `node_modules`")
	addCmd.Flags().Bool("save-build", false, "Install crate: packages as Cargo build dependencies")
	addCmd.Flags().Bool("save-catalog", false, "Save the new dependency to the default catalog. Shorthand for `--save-catalog-name=default`")
	addCmd.Flags().String("save-catalog-name", "", "Save the new dependency to the named catalog `<name>`")
	addCmd.Flags().BoolP("save-dev", "D", false, "Install the specified packages as devDependencies")
	addCmd.Flags().BoolP("save-exact", "E", false, "Saved dependencies will be configured with an exact version rather than using the default semver range operator")
	addCmd.Flags().BoolP("save-optional", "O", false, "Install the specified packages as optionalDependencies")
	addCmd.Flags().Bool("save-peer", false, "Using --save-peer will add one or more packages to peerDependencies and install them as dev dependencies")
	addCmd.Flags().String("save-prefix", "", "The prefix of the saved version range: `^` (default), `~`, `=` for an explicit exact pin, or empty for a bare exact version")
	addCmd.Flags().BoolP("save-prod", "P", false, "Install the specified packages as regular dependencies")
	addCmd.Flags().Bool("save-types", false, "Add available `@types` packages to `devDependencies` for packages without bundled types")
	addCmd.Flags().Bool("tilde", false, "Save the resolved version with a `~` range prefix. Equivalent to `--save-prefix=~`")
	addCmd.Flags().Bool("workspace", false, "Only add the dependency if a workspace project provides it. The dependency is saved under the `workspace:` protocol and linked to that project")
	addCmd.Flag("no-ignore-workspace-root-check").Hidden = true
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"cpu":          carapace.ActionValues("arm", "arm64", "ia32", "loong64", "mips", "mipsel", "ppc64", "riscv64", "s390", "s390x", "x64"),
		"libc":         carapace.ActionValues("glibc", "musl"),
		"lockfile-dir": carapace.ActionDirectories(),
		"os":           carapace.ActionValues("aix", "android", "darwin", "freebsd", "linux", "openbsd", "sunos", "win32"),
		"save-prefix":  carapace.ActionValues("^", "~", "="),
	})

	carapace.Gen(addCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.Value, "https://") {
				return git.ActionRepositorySearch(git.SearchOpts{}.Default())
			}

			return carapace.Batch(
				carapace.ActionFiles(),
				npm.ActionPackageSearch("").UnlessF(condition.CompletingPath),
				git.ActionRepositorySearch(git.SearchOpts{}.Default()).UnlessF(condition.CompletingPath),
			).ToA()
		}),
	)
}
