package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install packages",
	Aliases: []string{"i"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().StringSlice("cpu", nil, "CPU architectures whose platform-specific optional dependencies should be installed. Repeat or comma-separate for multiple values")
	installCmd.Flags().BoolP("dev", "D", false, "Install only devDependencies. Regular dependencies are skipped, and removed if already installed, regardless of `NODE_ENV`")
	installCmd.Flags().Bool("dry-run", false, "Show what an install would change without writing anything to disk")
	installCmd.Flags().String("fetch-min-speed-ki-bps", "", "Warn when a tarball download's average speed is below this many KiB/s")
	installCmd.Flags().String("fetch-timeout", "", "Per-request network timeout, in milliseconds")
	installCmd.Flags().String("fetch-warn-timeout-ms", "", "Warn when a registry metadata request takes longer than this many milliseconds")
	installCmd.Flags().Bool("fix-lockfile", false, "Repair broken lockfile entries by re-resolving their metadata while preserving compatible locked versions")
	installCmd.Flags().Bool("force", false, "Reinstall every package the lockfile names: relink packages an earlier install already materialized, and install optional dependencies whose `cpu` / `os` / `libc` / `engines` don't match the host instead of skipping them")
	installCmd.Flags().Bool("frozen-lockfile", false, "Don't generate a lockfile, and fail if an update to it is needed. This setting is enabled by default in CI when a lockfile is present")
	installCmd.Flags().Bool("frozen-store", false, "Open the store read-only and skip all store writes. For installing against a store on a read-only filesystem (e.g. a Nix store); pair with `--offline --frozen-lockfile`")
	installCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	installCmd.Flags().Bool("ignore-manifest-check", false, "Skip the check that `pnpm-lock.yaml` is up to date with `package.json` under `--frozen-lockfile`. For callers that just wrote the lockfile themselves and know the manifest is about to catch up")
	installCmd.Flags().Bool("ignore-pnpmfile", false, "Disable pnpm hooks defined in `.pnpmfile.cjs`, including the pnpmfiles of config dependencies")
	installCmd.Flags().Bool("ignore-scripts", false, "Don't run lifecycle scripts of the project or its dependencies. Packages are still installed; only their build scripts are skipped, and the install won't fail because of it")
	installCmd.Flags().StringSlice("libc", nil, "libc families whose platform-specific optional dependencies should be installed (`glibc`, `musl`). Repeat or comma-separate for multiple values")
	installCmd.Flags().String("lockfile-dir", "", "The directory in which `pnpm-lock.yaml` is created. Several projects may share a single lockfile")
	installCmd.Flags().Bool("lockfile-only", false, "Only update `pnpm-lock.yaml`. Don't download packages or write `node_modules`")
	installCmd.Flags().Bool("merge-git-branch-lockfiles", false, "Fold every per-branch lockfile (`pnpm-lock.<branch>.yaml`, written under the `gitBranchLockfile` setting) into `pnpm-lock.yaml` and delete them")
	installCmd.Flags().StringSlice("merge-git-branch-lockfiles-branch-pattern", nil, "Glob patterns naming the branches that merge the per-branch lockfiles, so a mainline branch does not have to pass `--merge-git-branch-lockfiles` by hand")
	installCmd.Flags().String("network-concurrency", "", "Maximum number of concurrent network requests during install")
	installCmd.Flags().Bool("no-frozen-lockfile", false, "Allow the lockfile to be updated, overriding a `frozenLockfile: true` setting")
	installCmd.Flags().Bool("no-frozen-store", false, "Allow store writes even when the configuration enables the read-only store")
	installCmd.Flags().Bool("no-ignore-scripts", false, "Run lifecycle scripts even when the configuration disables them")
	installCmd.Flags().Bool("no-offline", false, "Allow network fetches even when the configuration enables offline mode")
	installCmd.Flags().Bool("no-optional", false, "Don't install optionalDependencies")
	installCmd.Flags().Bool("no-prefer-frozen-lockfile", false, "Always re-resolve against the registry instead of preferring the existing lockfile")
	installCmd.Flags().Bool("no-prefer-offline", false, "Don't prefer cached packages even when the configuration enables it")
	installCmd.Flags().Bool("no-runtime", false, "Don't install runtime dependencies (`node`, `deno`, `bun`). Their archives aren't fetched and their bins aren't linked; the rest of the install proceeds normally")
	installCmd.Flags().Bool("no-trust-lockfile", false, "Verify the lockfile against supply-chain policies even when the configuration trusts it")
	installCmd.Flags().String("node-linker", "", "Which node linker to use: `isolated` (the default, a symlinked store), `hoisted` (a flat `node_modules`), or `pnp` (Plug'n'Play). Overrides the configured value")
	installCmd.Flags().Bool("offline", false, "Fail on a cache miss instead of fetching from the registry, using only packages already in the store")
	installCmd.Flags().Bool("optional", false, "Include optionalDependencies even when the configured default excludes them")
	installCmd.Flags().StringSlice("os", nil, "Operating systems whose platform-specific optional dependencies should be installed. Repeat or comma-separate for multiple values")
	installCmd.Flags().String("pnpr-server", "", "URL of a pnpr server to offload resolution and file fetching to. `node_modules` is still linked locally from the server-produced lockfile")
	installCmd.Flags().Bool("prefer-frozen-lockfile", false, "Prefer the existing lockfile over re-resolving, even when the manifest may have changed")
	installCmd.Flags().Bool("prefer-offline", false, "Prefer packages already in the cache over the network, even past their freshness window")
	installCmd.Flags().BoolP("prod", "P", false, "Install only production dependencies. devDependencies are skipped, and removed if already installed. Takes precedence over `NODE_ENV`")
	installCmd.Flags().Bool("production", false, "Install only production dependencies. devDependencies are skipped, and removed if already installed. Takes precedence over `NODE_ENV`")
	installCmd.Flags().Bool("trust-lockfile", false, "Skip verifying the lockfile against supply-chain policies")
	installCmd.Flags().Bool("update-checksums", false, "Refresh the integrity checksums in `pnpm-lock.yaml` from the registry. Cannot be combined with `--frozen-lockfile`")
	installCmd.Flags().String("user-agent", "", "`User-Agent` header to send on registry requests")
	installCmd.Flags().Bool("verify-deps-before-run-install", false, "Run the install already requested by `verifyDepsBeforeRun` without independently short-circuiting it as up to date")
	installCmd.Flag("verify-deps-before-run-install").Hidden = true

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"lockfile-dir": carapace.ActionDirectories(),
		"node-linker":  carapace.ActionValues("isolated", "hoisted", "pnp"),
	})

	rootCmd.AddCommand(installCmd)
}
