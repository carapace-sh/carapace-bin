package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ciCmd = &cobra.Command{
	Use:     "ci",
	Short:   "Runs clean then install with a frozen lockfile",
	Aliases: []string{"clean-install", "ic", "install-clean"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ciCmd).Standalone()

	ciCmd.Flags().StringSlice("cpu", nil, "CPU architectures whose platform-specific optional dependencies should be installed. Repeat or comma-separate for multiple values")
	ciCmd.Flags().BoolP("dev", "D", false, "Install only devDependencies. Regular dependencies are skipped, and removed if already installed, regardless of `NODE_ENV`")
	ciCmd.Flags().Bool("dry-run", false, "Show what an install would change without writing anything to disk")
	ciCmd.Flags().String("fetch-min-speed-ki-bps", "", "Warn when a tarball download's average speed is below this many KiB/s")
	ciCmd.Flags().String("fetch-timeout", "", "Per-request network timeout, in milliseconds")
	ciCmd.Flags().String("fetch-warn-timeout-ms", "", "Warn when a registry metadata request takes longer than this many milliseconds")
	ciCmd.Flags().Bool("fix-lockfile", false, "Repair broken lockfile entries by re-resolving their metadata while preserving compatible locked versions")
	ciCmd.Flags().Bool("force", false, "Reinstall every package the lockfile names: relink packages an earlier install already materialized, and install optional dependencies whose `cpu` / `os` / `libc` / `engines` don't match the host instead of skipping them")
	ciCmd.Flags().Bool("frozen-lockfile", false, "Don't generate a lockfile, and fail if an update to it is needed. This setting is enabled by default in CI when a lockfile is present")
	ciCmd.Flags().Bool("frozen-store", false, "Open the store read-only and skip all store writes. For installing against a store on a read-only filesystem (e.g. a Nix store); pair with `--offline --frozen-lockfile`")
	ciCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	ciCmd.Flags().Bool("ignore-manifest-check", false, "Skip the check that `pnpm-lock.yaml` is up to date with `package.json` under `--frozen-lockfile`. For callers that just wrote the lockfile themselves and know the manifest is about to catch up")
	ciCmd.Flags().Bool("ignore-pnpmfile", false, "Disable pnpm hooks defined in `.pnpmfile.cjs`, including the pnpmfiles of config dependencies")
	ciCmd.Flags().Bool("ignore-scripts", false, "Don't run lifecycle scripts of the project or its dependencies. Packages are still installed; only their build scripts are skipped, and the install won't fail because of it")
	ciCmd.Flags().StringSlice("libc", nil, "libc families whose platform-specific optional dependencies should be installed (`glibc`, `musl`). Repeat or comma-separate for multiple values")
	ciCmd.Flags().BoolP("lockfile", "l", false, "Also remove `pnpm-lock.yaml` files")
	ciCmd.Flags().String("lockfile-dir", "", "The directory in which `pnpm-lock.yaml` is created. Several projects may share a single lockfile")
	ciCmd.Flags().Bool("lockfile-only", false, "Only update `pnpm-lock.yaml`. Don't download packages or write `node_modules`")
	ciCmd.Flags().Bool("merge-git-branch-lockfiles", false, "Fold every per-branch lockfile (`pnpm-lock.<branch>.yaml`, written under the `gitBranchLockfile` setting) into `pnpm-lock.yaml` and delete them")
	ciCmd.Flags().StringSlice("merge-git-branch-lockfiles-branch-pattern", nil, "Glob patterns naming the branches that merge the per-branch lockfiles, so a mainline branch does not have to pass `--merge-git-branch-lockfiles` by hand")
	ciCmd.Flags().String("network-concurrency", "", "Maximum number of concurrent network requests during install")
	ciCmd.Flags().Bool("no-frozen-lockfile", false, "Allow the lockfile to be updated, overriding a `frozenLockfile: true` setting")
	ciCmd.Flags().Bool("no-frozen-store", false, "Allow store writes even when the configuration enables the read-only store")
	ciCmd.Flags().Bool("no-ignore-scripts", false, "Run lifecycle scripts even when the configuration disables them")
	ciCmd.Flags().Bool("no-offline", false, "Allow network fetches even when the configuration enables offline mode")
	ciCmd.Flags().Bool("no-optional", false, "Don't install optionalDependencies")
	ciCmd.Flags().Bool("no-prefer-frozen-lockfile", false, "Always re-resolve against the registry instead of preferring the existing lockfile")
	ciCmd.Flags().Bool("no-prefer-offline", false, "Don't prefer cached packages even when the configuration enables it")
	ciCmd.Flags().Bool("no-runtime", false, "Don't install runtime dependencies (`node`, `deno`, `bun`). Their archives aren't fetched and their bins aren't linked; the rest of the install proceeds normally")
	ciCmd.Flags().Bool("no-trust-lockfile", false, "Verify the lockfile against supply-chain policies even when the configuration trusts it")
	ciCmd.Flags().String("node-linker", "", "Which node linker to use: `isolated` (the default, a symlinked store), `hoisted` (a flat `node_modules`), or `pnp` (Plug'n'Play). Overrides the configured value")
	ciCmd.Flags().Bool("offline", false, "Fail on a cache miss instead of fetching from the registry, using only packages already in the store")
	ciCmd.Flags().Bool("optional", false, "Include optionalDependencies even when the configured default excludes them")
	ciCmd.Flags().StringSlice("os", nil, "Operating systems whose platform-specific optional dependencies should be installed. Repeat or comma-separate for multiple values")
	ciCmd.Flags().String("pnpr-server", "", "URL of a pnpr server to offload resolution and file fetching to. `node_modules` is still linked locally from the server-produced lockfile")
	ciCmd.Flags().Bool("prefer-frozen-lockfile", false, "Prefer the existing lockfile over re-resolving, even when the manifest may have changed")
	ciCmd.Flags().Bool("prefer-offline", false, "Prefer packages already in the cache over the network, even past their freshness window")
	ciCmd.Flags().BoolP("prod", "P", false, "Install only production dependencies. devDependencies are skipped, and removed if already installed. Takes precedence over `NODE_ENV`")
	ciCmd.Flags().Bool("production", false, "Install only production dependencies. devDependencies are skipped, and removed if already installed. Takes precedence over `NODE_ENV`")
	ciCmd.Flags().Bool("trust-lockfile", false, "Skip verifying the lockfile against supply-chain policies")
	ciCmd.Flags().Bool("update-checksums", false, "Refresh the integrity checksums in `pnpm-lock.yaml` from the registry. Cannot be combined with `--frozen-lockfile`")
	ciCmd.Flags().String("user-agent", "", "`User-Agent` header to send on registry requests")
	ciCmd.Flags().Bool("verify-deps-before-run-install", false, "Run the install already requested by `verifyDepsBeforeRun` without independently short-circuiting it as up to date")
	ciCmd.Flag("verify-deps-before-run-install").Hidden = true
	rootCmd.AddCommand(ciCmd)
}
