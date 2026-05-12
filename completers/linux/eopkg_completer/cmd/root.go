package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eopkg",
	Short: "Solus package manager",
	Long:  "https://github.com/solus-project/package-management/blob/master/man/eopkg.1.md",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	addGlobalFlags(rootCmd)

	rootCmd.AddCommand(
		command("add-repo <repo-name> <repo URI>", "Add a new repository", []string{"ar"}, addRepoFlags),
		command("autoremove <package>...", "Remove packages and unneeded reverse dependencies", []string{"rmf"}, autoremoveFlags, eopkg.ActionInstalledPackages()),
		command("blame <package>", "Show package changelog history", []string{"bl"}, blameFlags, eopkg.ActionAvailablePackages()),
		command("build <pspec.xml>", "Build a legacy eopkg package", []string{"bi"}, buildFlags, carapace.ActionFiles(".xml")),
		command("check [package]...", "Check installed package integrity", nil, checkFlags, eopkg.ActionInstalledPackages()),
		command("clean", "Delete stale eopkg lock files", nil, nil),
		command("configure-pending", "Configure pending packages", []string{"cp"}, configurePendingFlags),
		command("delete-cache", "Clear temporary eopkg caches", []string{"dc"}, nil),
		command("delta <oldpackage> <newpackage>", "Construct a delta package", []string{"dt"}, deltaFlags, carapace.ActionFiles(".eopkg")),
		command("disable-repo <name>", "Disable a repository", []string{"dr"}, nil, eopkg.ActionRepositories()),
		command("emerge <name>", "Build a legacy source package", []string{"em"}, buildFlags, eopkg.ActionAvailablePackages()),
		command("enable-repo <name>", "Enable a repository", []string{"er"}, nil, eopkg.ActionRepositories()),
		command("fetch <name>", "Download a package", []string{"fc"}, fetchFlags, eopkg.ActionAvailablePackages()),
		command("graph <package>", "Generate a package dependency graph", nil, graphFlags, eopkg.ActionAvailablePackages()),
		command("help [subcommand]", "Display help topics", []string{"?"}, nil, commandValues()),
		command("history", "Manage transaction history", []string{"hs"}, historyFlags),
		command("index <directory>", "Produce an eopkg repository index", []string{"ix"}, indexFlags, carapace.ActionDirectories()),
		command("info <package>", "Show package information", nil, infoFlags, packageOrFileAction()),
		command("install <name>", "Install a package or local .eopkg file", []string{"it"}, installFlags, packageOrFileAction()),
		command("list-available [repo]", "List available packages", []string{"la"}, listAvailableFlags, eopkg.ActionRepositories()),
		command("list-components", "List available package components", []string{"lc"}, listComponentsFlags),
		command("list-installed", "List installed packages", []string{"li"}, listInstalledFlags),
		command("list-newest [repo]", "List newest repository packages", []string{"ln"}, listNewestFlags, eopkg.ActionRepositories()),
		command("list-pending", "List packages requiring configuration", []string{"lp"}, nil),
		command("list-repo", "List configured repositories", []string{"lr"}, nil),
		command("list-sources", "List legacy source packages", []string{"ls"}, listSourcesFlags),
		command("list-upgrades", "List available package upgrades", []string{"lu"}, listUpgradesFlags),
		command("rebuild-db", "Rebuild eopkg databases", []string{"rdb"}, rebuildDBFlags),
		command("remove <package>...", "Remove installed packages", []string{"rm"}, removeFlags, eopkg.ActionInstalledPackages()),
		command("remove-orphans", "Remove unneeded automatically installed packages", []string{"rmo"}, removeOrphansFlags),
		command("remove-repo <name>", "Remove a repository", []string{"rr"}, nil, eopkg.ActionRepositories()),
		command("search <term>", "Search packages", []string{"sr"}, searchFlags),
		command("search-file <path>", "Find package owning a local path", []string{"sf"}, searchFileFlags, carapace.ActionFiles()),
		command("update-repo [name]", "Update repository indexes", []string{"ur"}, updateRepoFlags, eopkg.ActionRepositories()),
		command("upgrade [package]...", "Upgrade the system or selected packages", []string{"up"}, upgradeFlags, eopkg.ActionAvailablePackages()),
	)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"destdir": carapace.ActionDirectories(),
	})
}

func command(use string, short string, aliases []string, setup func(*cobra.Command), positional ...carapace.Action) *cobra.Command {
	cmd := &cobra.Command{
		Use:     use,
		Short:   short,
		Aliases: aliases,
		Run:     func(cmd *cobra.Command, args []string) {},
	}
	carapace.Gen(cmd).Standalone()
	if setup != nil {
		setup(cmd)
	}
	if len(positional) > 0 {
		carapace.Gen(cmd).PositionalAnyCompletion(positional[0])
	}
	return cmd
}

func addGlobalFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("bandwidth-limit", "L", "", "Limit bandwidth usage in KB/s")
	cmd.Flags().BoolP("debug", "d", false, "Enable debug information and backtraces")
	cmd.Flags().StringP("destdir", "D", "", "Change the system root")
	cmd.Flags().BoolP("help", "h", false, "Show help")
	cmd.Flags().BoolP("no-color", "N", false, "Disable colored output")
	cmd.Flags().StringP("password", "p", "", "Password for Basic-Auth repositories")
	cmd.Flags().StringP("username", "u", "", "Username for Basic-Auth repositories")
	cmd.Flags().BoolP("verbose", "v", false, "Show detailed output")
	cmd.Flags().Bool("version", false, "Show version")
	cmd.Flags().BoolP("yes-all", "y", false, "Assume yes in all yes/no queries")
}

func commandValues() carapace.Action {
	return carapace.ActionValuesDescribed(
		"add-repo", "add a new repository",
		"autoremove", "remove packages and unneeded reverse dependencies",
		"blame", "show package changelog history",
		"build", "build a legacy eopkg package",
		"check", "check installed package integrity",
		"clean", "delete stale eopkg lock files",
		"configure-pending", "configure pending packages",
		"delete-cache", "clear temporary eopkg caches",
		"delta", "construct a delta package",
		"disable-repo", "disable a repository",
		"emerge", "build a legacy source package",
		"enable-repo", "enable a repository",
		"fetch", "download a package",
		"graph", "generate a package dependency graph",
		"history", "manage transaction history",
		"index", "produce a repository index",
		"info", "show package information",
		"install", "install a package",
		"list-available", "list available packages",
		"list-components", "list available package components",
		"list-installed", "list installed packages",
		"list-newest", "list newest repository packages",
		"list-pending", "list packages requiring configuration",
		"list-repo", "list configured repositories",
		"list-sources", "list legacy source packages",
		"list-upgrades", "list available package upgrades",
		"rebuild-db", "rebuild databases",
		"remove", "remove installed packages",
		"remove-orphans", "remove orphaned packages",
		"remove-repo", "remove a repository",
		"search", "search packages",
		"search-file", "find the package owning a local path",
		"update-repo", "update repository indexes",
		"upgrade", "upgrade packages",
	).Tag("commands")
}

func packageOrFileAction() carapace.Action {
	return carapace.Batch(
		carapace.ActionFiles(".eopkg"),
		eopkg.ActionAvailablePackages().UnlessF(condition.CompletingPath),
	).ToA()
}

func addRepoFlags(cmd *cobra.Command) {
	cmd.Flags().String("at", "", "Insert the repository at the given index")
	cmd.Flags().Bool("ignore-check", false, "Ignore metadata distribution checks")
	cmd.Flags().Bool("no-fetch", false, "Register the repository without downloading the index")
}

func autoremoveFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("dry-run", "n", false, "Show what would happen without making changes")
	cmd.Flags().Bool("ignore-comar", false, "Bypass system configuration")
	cmd.Flags().Bool("ignore-dependency", false, "Ignore reverse dependency validation")
	cmd.Flags().Bool("ignore-safety", false, "Ignore the system.base safety switch")
	cmd.Flags().BoolP("purge", "p", false, "Remove configuration files too")
}

func removeFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("component", "c", "", "Operate on a component")
	cmd.Flags().BoolP("dry-run", "n", false, "Show what would happen without making changes")
	cmd.Flags().Bool("ignore-comar", false, "Bypass system configuration")
	cmd.Flags().Bool("ignore-dependency", false, "Ignore reverse dependency validation")
	cmd.Flags().Bool("ignore-safety", false, "Ignore the system.base safety switch")
	cmd.Flags().BoolP("purge", "p", false, "Remove configuration files too")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"component": eopkg.ActionComponents(),
	})
}

func removeOrphansFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("dry-run", "n", false, "Show what would happen without making changes")
	cmd.Flags().Bool("ignore-comar", false, "Bypass system configuration")
	cmd.Flags().Bool("ignore-safety", false, "Ignore the system.base safety switch")
	cmd.Flags().BoolP("purge", "p", false, "Remove configuration files too")
}

func blameFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("all", "a", false, "Show the entire package history")
	cmd.Flags().StringP("release", "r", "", "Show blame for the given release")
}

func buildFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("build", false, "Build package")
	cmd.Flags().Bool("check", false, "Run checks")
	cmd.Flags().StringP("component", "c", "", "Operate on a component")
	cmd.Flags().Bool("create-static", false, "Create a static package")
	cmd.Flags().Bool("fetch", false, "Fetch source archives")
	cmd.Flags().Bool("ignore-action-errors", false, "Ignore action errors")
	cmd.Flags().Bool("ignore-build-no", false, "Ignore build number")
	cmd.Flags().Bool("ignore-check", false, "Ignore package metadata checks")
	cmd.Flags().Bool("ignore-dependency", false, "Ignore dependency validation")
	cmd.Flags().Bool("ignore-file-conflicts", false, "Ignore file conflicts")
	cmd.Flags().Bool("ignore-package-conflicts", false, "Ignore package conflicts")
	cmd.Flags().Bool("ignore-safety", false, "Ignore safety checks")
	cmd.Flags().Bool("ignore-sandbox", false, "Ignore sandboxing")
	cmd.Flags().Bool("install", false, "Install into the package image")
	cmd.Flags().StringP("output-dir", "O", "", "Write output to directory")
	cmd.Flags().Bool("package", false, "Build the binary package")
	cmd.Flags().StringP("package-format", "F", "", "Override package format")
	cmd.Flags().Bool("quiet", false, "Reduce output")
	cmd.Flags().Bool("setup", false, "Run setup")
	cmd.Flags().Bool("unpack", false, "Unpack sources")
	cmd.Flags().Bool("use-quilt", false, "Use quilt")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"component":  eopkg.ActionComponents(),
		"output-dir": carapace.ActionDirectories(),
	})
}

func checkFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("component", "c", "", "Check packages under a component")
	cmd.Flags().Bool("config", false, "Only check configuration files")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"component": eopkg.ActionComponents(),
	})
}

func configurePendingFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("dry-run", "n", false, "Show what would happen without making changes")
	cmd.Flags().Bool("ignore-comar", false, "Bypass system configuration")
	cmd.Flags().Bool("ignore-dependency", false, "Ignore dependency validation")
	cmd.Flags().Bool("ignore-safety", false, "Ignore the system.base safety switch")
}

func deltaFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("newest-package", "t", "", "Override the new package detection")
	cmd.Flags().StringP("output-dir", "O", "", "Write output to directory")
	cmd.Flags().StringP("package-format", "F", "", "Override package format")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"newest-package": carapace.ActionFiles(".eopkg"),
		"output-dir":     carapace.ActionDirectories(),
	})
}

func fetchFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("output-dir", "o", "", "Write downloaded package to directory")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"output-dir": carapace.ActionDirectories(),
	})
}

func graphFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("ignore-installed", false, "Ignore installed packages")
	cmd.Flags().Bool("installed", false, "Use installed packages")
	cmd.Flags().StringP("output", "o", "", "Write graph output to file")
	cmd.Flags().StringP("repository", "r", "", "Only inspect a repository")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"output":     carapace.ActionFiles(),
		"repository": eopkg.ActionRepositories(),
	})
}

func historyFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("last", "l", "", "Only output the last operations")
	cmd.Flags().BoolP("snapshot", "s", false, "Create a transaction snapshot")
	cmd.Flags().StringP("takeback", "t", "", "Roll back to a transaction")
}

func indexFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("absolute-urls", "a", false, "Use absolute URLs in the index")
	cmd.Flags().String("compression-types", "", "Comma-separated compression types")
	cmd.Flags().StringP("output", "o", "", "Override output file")
	cmd.Flags().Bool("skip-signing", false, "Do not GPG sign the index")
	cmd.Flags().Bool("skip-sources", false, "Do not include legacy pspec.xml sources")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"compression-types": carapace.ActionValues("bz2", "xz").UniqueList(","),
		"output":            carapace.ActionFiles(),
	})
}

func infoFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("component", "c", "", "Show information about a component")
	cmd.Flags().BoolP("files", "f", false, "Show package files")
	cmd.Flags().BoolP("files-path", "F", false, "Only show package files")
	cmd.Flags().BoolP("short", "s", false, "Use compact output")
	cmd.Flags().Bool("xml", false, "Emit original XML metadata")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"component": eopkg.ActionComponents(),
	})
}

func installFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("component", "c", "", "Install an entire component")
	cmd.Flags().BoolP("dry-run", "n", false, "Show what would happen without making changes")
	cmd.Flags().StringP("exclude", "x", "", "Exclude packages or components")
	cmd.Flags().String("exclude-from", "", "Read excluded packages or components from file")
	cmd.Flags().BoolP("fetch-only", "f", false, "Only download packages")
	cmd.Flags().Bool("ignore-build-no", false, "Ignore build number")
	cmd.Flags().Bool("ignore-check", false, "Ignore package metadata checks")
	cmd.Flags().Bool("ignore-comar", false, "Bypass system configuration")
	cmd.Flags().Bool("ignore-dependency", false, "Ignore dependency validation")
	cmd.Flags().Bool("ignore-file-conflicts", false, "Ignore file conflicts")
	cmd.Flags().Bool("ignore-package-conflicts", false, "Ignore package conflicts")
	cmd.Flags().Bool("ignore-safety", false, "Ignore the system.base safety switch")
	cmd.Flags().Bool("reinstall", false, "Reinstall already installed packages")
	cmd.Flags().StringP("repository", "r", "", "Use a specific repository")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"component":    eopkg.ActionComponents(),
		"exclude":      eopkg.ActionAvailablePackages(),
		"exclude-from": carapace.ActionFiles(),
		"repository":   eopkg.ActionRepositories(),
	})
}

func listAvailableFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("component", "c", "", "List packages under a component")
	cmd.Flags().BoolP("long", "l", false, "Use long output")
	cmd.Flags().BoolP("uninstalled", "U", false, "Only show uninstalled packages")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"component": eopkg.ActionComponents(),
	})
}

func listComponentsFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("long", "l", false, "Show full component details")
	cmd.Flags().StringP("repository", "r", "", "Only list components in a repository")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"repository": eopkg.ActionRepositories(),
	})
}

func listInstalledFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("automatic", "a", false, "Show automatically installed packages")
	cmd.Flags().StringP("build-host", "b", "", "Only show packages from a build host")
	cmd.Flags().StringP("component", "c", "", "Only show packages from a component")
	cmd.Flags().BoolP("install-info", "i", false, "Show installation information")
	cmd.Flags().BoolP("long", "l", false, "Show full details")
	cmd.Flags().Bool("without-buildno", false, "Hide build numbers")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"component": eopkg.ActionComponents(),
	})
}

func listNewestFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("last", "l", "", "Only show packages since the nth repository update")
	cmd.Flags().StringP("since", "s", "", "Show packages newer than YYYY-MM-DD")
}

func listSourcesFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("long", "l", false, "Show detailed source package information")
}

func listUpgradesFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("component", "c", "", "Only show upgrades from a component")
	cmd.Flags().Bool("ignore-build-no", false, "Ignore build number")
	cmd.Flags().BoolP("install-info", "i", false, "Show installation information")
	cmd.Flags().BoolP("long", "l", false, "Show detailed package information")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"component": eopkg.ActionComponents(),
	})
}

func rebuildDBFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("files", "f", false, "Only rebuild the files database")
}

func searchFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("description", false, "Only search package descriptions")
	cmd.Flags().BoolP("installdb", "i", false, "Only search installed packages")
	cmd.Flags().StringP("language", "l", "", "Only search summaries/descriptions for language code")
	cmd.Flags().Bool("name", false, "Only search package names")
	cmd.Flags().StringP("repository", "r", "", "Only search within a repository")
	cmd.Flags().BoolP("sourcedb", "s", false, "Only search source repositories")
	cmd.Flags().Bool("summary", false, "Only search package summaries")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"language":   carapace.ActionValues("en"),
		"repository": eopkg.ActionRepositories(),
	})
}

func searchFileFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("long", "l", false, "Show detailed information")
	cmd.Flags().BoolP("quiet", "q", false, "Only show package names")
}

func updateRepoFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("force", "f", false, "Force repository index update")
}

func upgradeFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("bypass-update-repo", "b", false, "Do not update repositories first")
	cmd.Flags().StringP("component", "c", "", "Only upgrade packages from a component")
	cmd.Flags().BoolP("dry-run", "n", false, "Show what would happen without making changes")
	cmd.Flags().StringP("exclude", "x", "", "Exclude packages or components")
	cmd.Flags().String("exclude-from", "", "Read excluded packages or components from file")
	cmd.Flags().BoolP("fetch-only", "f", false, "Only download packages")
	cmd.Flags().Bool("ignore-build-no", false, "Ignore build number")
	cmd.Flags().Bool("ignore-comar", false, "Bypass system configuration")
	cmd.Flags().Bool("ignore-dependency", false, "Ignore dependency validation")
	cmd.Flags().Bool("ignore-file-conflicts", false, "Ignore file conflicts")
	cmd.Flags().Bool("ignore-package-conflicts", false, "Ignore package conflicts")
	cmd.Flags().Bool("ignore-safety", false, "Ignore the system.base safety switch")
	cmd.Flags().StringP("repository", "r", "", "Only upgrade packages from a repository")
	cmd.Flags().Bool("security-only", false, "Only apply security updates")
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"component":    eopkg.ActionComponents(),
		"exclude":      eopkg.ActionAvailablePackages(),
		"exclude-from": carapace.ActionFiles(),
		"repository":   eopkg.ActionRepositories(),
	})
}
