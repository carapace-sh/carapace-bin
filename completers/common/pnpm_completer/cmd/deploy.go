package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a package from a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployCmd).Standalone()

	deployCmd.Flags().StringSlice("allow-build", nil, "Package names allowed to run lifecycle (build) scripts during this install, appended to `allowBuilds`. Prefix a name with `!` to deny its scripts instead. May be repeated")
	deployCmd.Flags().Bool("auto-dedupe", false, "Deduplicate compatible dependency versions during installation")
	deployCmd.Flags().StringSlice("cpu", nil, "CPU architectures whose platform-specific optional dependencies should be installed. Repeat or comma-separate for multiple values")
	deployCmd.Flags().BoolP("dev", "D", false, "Install only devDependencies. Regular dependencies are skipped, and removed if already installed")
	deployCmd.Flags().Bool("dry-run", false, "Show what an install would change without writing anything to disk")
	deployCmd.Flags().String("fetch-min-speed-ki-bps", "", "Warn when a tarball download's average speed is below this many KiB/s")
	deployCmd.Flags().String("fetch-timeout", "", "Per-request network timeout, in milliseconds")
	deployCmd.Flags().String("fetch-warn-timeout-ms", "", "Warn when a registry metadata request takes longer than this many milliseconds")
	deployCmd.Flags().Bool("fix-lockfile", false, "Repair broken lockfile entries by re-resolving their metadata while preserving compatible locked versions")
	deployCmd.Flags().Bool("force", false, "Reinstall every package the lockfile names: relink packages an earlier install already materialized, and install optional dependencies whose `cpu` / `os` / `libc` / `engines` don't match the host instead of skipping them")
	deployCmd.Flags().Bool("frozen-lockfile", false, "Don't generate a lockfile, and fail if an update to it is needed. This setting is enabled by default in CI when a lockfile is present")
	deployCmd.Flags().Bool("frozen-store", false, "Open the store read-only and skip all store writes. For installing against a store on a read-only filesystem (e.g. a Nix store); pair with `--offline --frozen-lockfile`")
	deployCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	deployCmd.Flags().Bool("ignore-manifest-check", false, "Skip the check that `pnpm-lock.yaml` is up to date with `package.json` under `--frozen-lockfile`. For callers that just wrote the lockfile themselves and know the manifest is about to catch up")
	deployCmd.Flags().Bool("ignore-pnpmfile", false, "Disable pnpm hooks defined in `.pnpmfile.cjs`, including the pnpmfiles of config dependencies")
	deployCmd.Flags().Bool("ignore-scripts", false, "Don't run lifecycle scripts of the project or its dependencies. Packages are still installed; only their build scripts are skipped, and the install won't fail because of it")
	deployCmd.Flags().Bool("legacy", false, "Use the legacy deploy implementation")
	deployCmd.Flags().StringSlice("libc", nil, "libc families whose platform-specific optional dependencies should be installed (`glibc`, `musl`). Repeat or comma-separate for multiple values")
	deployCmd.Flags().String("lockfile-dir", "", "The directory in which `pnpm-lock.yaml` is created. Several projects may share a single lockfile")
	deployCmd.Flags().Bool("lockfile-only", false, "Only update `pnpm-lock.yaml`. Don't download packages or write `node_modules`")
	deployCmd.Flags().Bool("merge-git-branch-lockfiles", false, "Fold every per-branch lockfile (`pnpm-lock.<branch>.yaml`, written under the `gitBranchLockfile` setting) into `pnpm-lock.yaml` and delete them")
	deployCmd.Flags().StringSlice("merge-git-branch-lockfiles-branch-pattern", nil, "Glob patterns naming the branches that merge the per-branch lockfiles, so a mainline branch does not have to pass `--merge-git-branch-lockfiles` by hand")
	deployCmd.Flags().String("network-concurrency", "", "Maximum number of concurrent network requests during install")
	deployCmd.Flags().Bool("no-auto-dedupe", false, "Disable automatic deduplication configured in pnpm-workspace.yaml")
	deployCmd.Flags().Bool("no-frozen-lockfile", false, "Allow the lockfile to be updated, overriding a `frozenLockfile: true` setting")
	deployCmd.Flags().Bool("no-frozen-store", false, "Allow store writes even when the configuration enables the read-only store")
	deployCmd.Flags().Bool("no-ignore-scripts", false, "Run lifecycle scripts even when the configuration disables them")
	deployCmd.Flags().Bool("no-offline", false, "Allow network fetches even when the configuration enables offline mode")
	deployCmd.Flags().Bool("no-optional", false, "Don't install optionalDependencies")
	deployCmd.Flags().Bool("no-prefer-frozen-lockfile", false, "Always re-resolve against the registry instead of preferring the existing lockfile")
	deployCmd.Flags().Bool("no-prefer-offline", false, "Don't prefer cached packages even when the configuration enables it")
	deployCmd.Flags().Bool("no-runtime", false, "Don't install runtime dependencies (`node`, `deno`, `bun`). Their archives aren't fetched and their bins aren't linked; the rest of the install proceeds normally")
	deployCmd.Flags().Bool("no-trust-lockfile", false, "Verify the lockfile against supply-chain policies even when the configuration trusts it")
	deployCmd.Flags().String("node-linker", "", "Which node linker to use: `isolated` (the default, a symlinked store), `hoisted` (a flat `node_modules`), or `pnp` (Plug'n'Play). Overrides the configured value")
	deployCmd.Flags().Bool("offline", false, "Fail on a cache miss instead of fetching from the registry, using only packages already in the store")
	deployCmd.Flags().Bool("optional", false, "Include optionalDependencies even when the configured default excludes them")
	deployCmd.Flags().StringSlice("os", nil, "Operating systems whose platform-specific optional dependencies should be installed. Repeat or comma-separate for multiple values")
	deployCmd.Flags().String("pnpr-server", "", "URL of a pnpr server to offload resolution and file fetching to. `node_modules` is still linked locally from the server-produced lockfile")
	deployCmd.Flags().Bool("prefer-frozen-lockfile", false, "Prefer the existing lockfile over re-resolving, even when the manifest may have changed")
	deployCmd.Flags().Bool("prefer-offline", false, "Prefer packages already in the cache over the network, even past their freshness window")
	deployCmd.Flags().BoolP("prod", "P", false, "Install only production dependencies. devDependencies are skipped, and removed if already installed")
	deployCmd.Flags().Bool("production", false, "Install only production dependencies. devDependencies are skipped, and removed if already installed")
	deployCmd.Flags().Bool("trust-lockfile", false, "Skip verifying the lockfile against supply-chain policies")
	deployCmd.Flags().Bool("update-checksums", false, "Refresh the integrity checksums in `pnpm-lock.yaml` from the registry. Cannot be combined with `--frozen-lockfile`")
	deployCmd.Flags().String("user-agent", "", "`User-Agent` header to send on registry requests")
	deployCmd.Flags().Bool("verify-deps-before-run-install", false, "Run the install already requested by `verifyDepsBeforeRun` without independently short-circuiting it as up to date")
	deployCmd.Flag("verify-deps-before-run-install").Hidden = true
	rootCmd.AddCommand(deployCmd)

	carapace.Gen(deployCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(deployCmd).FlagCompletion(carapace.ActionMap{
		"cpu":         carapace.ActionValues("arm", "arm64", "ia32", "loong64", "mips", "mipsel", "ppc64", "riscv64", "s390", "s390x", "x64"),
		"libc":        carapace.ActionValues("glibc", "musl"),
		"node-linker": carapace.ActionValues("isolated", "hoisted", "pnp"),
		"os":          carapace.ActionValues("aix", "android", "darwin", "freebsd", "linux", "openbsd", "sunos", "win32"),
	})
}
