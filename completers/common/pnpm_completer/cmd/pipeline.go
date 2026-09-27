package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Runs a named pipeline of workspace tasks the way a CI run would: a frozen install, affected-since-base selection, the task graph in dependency order without bailing, and cached task results restored instead of re-run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipelineCmd).Standalone()

	pipelineCmd.Flags().StringSlice("allow-build", nil, "Package names allowed to run lifecycle (build) scripts during this install, appended to `allowBuilds`. Prefix a name with `!` to deny its scripts instead. May be repeated")
	pipelineCmd.Flags().Bool("auto-dedupe", false, "Deduplicate compatible dependency versions during installation")
	pipelineCmd.Flags().String("base", "", "The git ref the affected selection diffs against (its merge base with HEAD). Overrides the `pipelineBase` setting")
	pipelineCmd.Flags().String("branch", "main", "The branch the watch agent follows")
	pipelineCmd.Flags().StringSlice("cpu", nil, "CPU architectures whose platform-specific optional dependencies should be installed. Repeat or comma-separate for multiple values")
	pipelineCmd.Flags().BoolP("dev", "D", false, "Install only devDependencies. Regular dependencies are skipped, and removed if already installed")
	pipelineCmd.Flags().Bool("dry-run", false, "Show what an install would change without writing anything to disk")
	pipelineCmd.Flags().String("fetch-min-speed-ki-bps", "", "Warn when a tarball download's average speed is below this many KiB/s")
	pipelineCmd.Flags().String("fetch-timeout", "", "Per-request network timeout, in milliseconds")
	pipelineCmd.Flags().String("fetch-warn-timeout-ms", "", "Warn when a registry metadata request takes longer than this many milliseconds")
	pipelineCmd.Flags().Bool("fix-lockfile", false, "Repair broken lockfile entries by re-resolving their metadata while preserving compatible locked versions")
	pipelineCmd.Flags().Bool("force", false, "Reinstall every package the lockfile names: relink packages an earlier install already materialized, and install optional dependencies whose `cpu` / `os` / `libc` / `engines` don't match the host instead of skipping them")
	pipelineCmd.Flags().Bool("frozen-lockfile", false, "Don't generate a lockfile, and fail if an update to it is needed. This setting is enabled by default in CI when a lockfile is present")
	pipelineCmd.Flags().Bool("frozen-store", false, "Open the store read-only and skip all store writes. For installing against a store on a read-only filesystem (e.g. a Nix store); pair with `--offline --frozen-lockfile`")
	pipelineCmd.Flags().Bool("full", false, "Run the pipeline over every workspace project instead of the affected-since-base selection")
	pipelineCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pipelineCmd.Flags().Bool("ignore-manifest-check", false, "Skip the check that `pnpm-lock.yaml` is up to date with `package.json` under `--frozen-lockfile`. For callers that just wrote the lockfile themselves and know the manifest is about to catch up")
	pipelineCmd.Flags().Bool("ignore-pnpmfile", false, "Disable pnpm hooks defined in `.pnpmfile.cjs`, including the pnpmfiles of config dependencies")
	pipelineCmd.Flags().Bool("ignore-scripts", false, "Don't run lifecycle scripts of the project or its dependencies. Packages are still installed; only their build scripts are skipped, and the install won't fail because of it")
	pipelineCmd.Flags().String("interval", "30", "Seconds between polls of the watched repository")
	pipelineCmd.Flags().Bool("json", false, "With `--dry-run`, print the tasks and their resolved dependency edges as JSON")
	pipelineCmd.Flags().StringSlice("libc", nil, "libc families whose platform-specific optional dependencies should be installed (`glibc`, `musl`). Repeat or comma-separate for multiple values")
	pipelineCmd.Flags().String("lockfile-dir", "", "The directory in which `pnpm-lock.yaml` is created. Several projects may share a single lockfile")
	pipelineCmd.Flags().Bool("lockfile-only", false, "Only update `pnpm-lock.yaml`. Don't download packages or write `node_modules`")
	pipelineCmd.Flags().Bool("merge-git-branch-lockfiles", false, "Fold every per-branch lockfile (`pnpm-lock.<branch>.yaml`, written under the `gitBranchLockfile` setting) into `pnpm-lock.yaml` and delete them")
	pipelineCmd.Flags().StringSlice("merge-git-branch-lockfiles-branch-pattern", nil, "Glob patterns naming the branches that merge the per-branch lockfiles, so a mainline branch does not have to pass `--merge-git-branch-lockfiles` by hand")
	pipelineCmd.Flags().String("network-concurrency", "", "Maximum number of concurrent network requests during install")
	pipelineCmd.Flags().Bool("no-auto-dedupe", false, "Disable automatic deduplication configured in pnpm-workspace.yaml")
	pipelineCmd.Flags().Bool("no-cache", false, "Run every task without reading or writing cached results or Cargo snapshots")
	pipelineCmd.Flags().Bool("no-frozen-lockfile", false, "Allow the lockfile to be updated, overriding a `frozenLockfile: true` setting")
	pipelineCmd.Flags().Bool("no-frozen-store", false, "Allow store writes even when the configuration enables the read-only store")
	pipelineCmd.Flags().Bool("no-ignore-scripts", false, "Run lifecycle scripts even when the configuration disables them")
	pipelineCmd.Flags().Bool("no-offline", false, "Allow network fetches even when the configuration enables offline mode")
	pipelineCmd.Flags().Bool("no-optional", false, "Don't install optionalDependencies")
	pipelineCmd.Flags().Bool("no-prefer-frozen-lockfile", false, "Always re-resolve against the registry instead of preferring the existing lockfile")
	pipelineCmd.Flags().Bool("no-prefer-offline", false, "Don't prefer cached packages even when the configuration enables it")
	pipelineCmd.Flags().Bool("no-runtime", false, "Don't install runtime dependencies (`node`, `deno`, `bun`). Their archives aren't fetched and their bins aren't linked; the rest of the install proceeds normally")
	pipelineCmd.Flags().Bool("no-trust-lockfile", false, "Verify the lockfile against supply-chain policies even when the configuration trusts it")
	pipelineCmd.Flags().String("node-linker", "", "Which node linker to use: `isolated` (the default, a symlinked store), `hoisted` (a flat `node_modules`), or `pnp` (Plug'n'Play). Overrides the configured value")
	pipelineCmd.Flags().Bool("offline", false, "Fail on a cache miss instead of fetching from the registry, using only packages already in the store")
	pipelineCmd.Flags().Bool("once", false, "With `--watch`: poll once, build if there is a new revision, and exit")
	pipelineCmd.Flags().Bool("optional", false, "Include optionalDependencies even when the configured default excludes them")
	pipelineCmd.Flags().StringSlice("os", nil, "Operating systems whose platform-specific optional dependencies should be installed. Repeat or comma-separate for multiple values")
	pipelineCmd.Flags().String("pnpr-server", "", "URL of a pnpr server to offload resolution and file fetching to. `node_modules` is still linked locally from the server-produced lockfile")
	pipelineCmd.Flags().Bool("prefer-frozen-lockfile", false, "Prefer the existing lockfile over re-resolving, even when the manifest may have changed")
	pipelineCmd.Flags().Bool("prefer-offline", false, "Prefer packages already in the cache over the network, even past their freshness window")
	pipelineCmd.Flags().BoolP("prod", "P", false, "Install only production dependencies. devDependencies are skipped, and removed if already installed")
	pipelineCmd.Flags().Bool("production", false, "Install only production dependencies. devDependencies are skipped, and removed if already installed")
	pipelineCmd.Flags().String("repo", "", "The repository the watch agent polls and builds — a URL or a local path, anything git accepts as a remote")
	pipelineCmd.Flags().Bool("report", false, "Publish the run's summary and event stream to the configured pnpr server (the `pnprServer` setting) once the run settles")
	pipelineCmd.Flags().String("report-to", "", "Publish the run to this pnpr server instead of the `pnprServer` setting — which also drives install offloading, so a server that only stores runs is better named here")
	pipelineCmd.Flags().Bool("trust-lockfile", false, "Skip verifying the lockfile against supply-chain policies")
	pipelineCmd.Flags().Bool("update-checksums", false, "Refresh the integrity checksums in `pnpm-lock.yaml` from the registry. Cannot be combined with `--frozen-lockfile`")
	pipelineCmd.Flags().String("user-agent", "", "`User-Agent` header to send on registry requests")
	pipelineCmd.Flags().Bool("verify-deps-before-run-install", false, "Run the install already requested by `verifyDepsBeforeRun` without independently short-circuiting it as up to date")
	pipelineCmd.Flags().Bool("watch", false, "Watch a git repository and run the pipeline for every new revision of a branch, instead of running once against the current directory")
	pipelineCmd.Flag("verify-deps-before-run-install").Hidden = true
	rootCmd.AddCommand(pipelineCmd)

	carapace.Gen(pipelineCmd).FlagCompletion(carapace.ActionMap{
		"cpu":         carapace.ActionValues("arm", "arm64", "ia32", "loong64", "mips", "mipsel", "ppc64", "riscv64", "s390", "s390x", "x64"),
		"libc":        carapace.ActionValues("glibc", "musl"),
		"node-linker": carapace.ActionValues("isolated", "hoisted", "pnp"),
		"os":          carapace.ActionValues("aix", "android", "darwin", "freebsd", "linux", "openbsd", "sunos", "win32"),
	})
}
