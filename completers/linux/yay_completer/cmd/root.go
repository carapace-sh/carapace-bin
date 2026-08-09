package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yay"
	"github.com/carapace-sh/carapace-bin/pkg/util/embed"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yay",
	Short: "An AUR Helper written in Go",
	Long:  "https://github.com/Jguer/yay",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("aur", "a", false, "Assume targets are from the AUR")
	rootCmd.Flags().String("aurrpcurl", "", "Set an alternative URL for the AUR /rpc endpoint")
	rootCmd.Flags().String("aururl", "", "Set an alternative AUR URL")
	rootCmd.Flags().String("builddir", "", "Directory used to download and run PKGBUILDS")
	rootCmd.Flags().String("completioninterval", "", "Time in days to refresh completion cache")
	rootCmd.Flags().String("config", "", "pacman.conf file to use")
	rootCmd.Flags().String("editor", "", "Editor to use when editing PKGBUILDs")
	rootCmd.Flags().String("editorflags", "", "Pass arguments to editor")
	rootCmd.Flags().String("git", "", "git command to use")
	rootCmd.Flags().String("gitflags", "", "Pass arguments to git")
	rootCmd.Flags().String("gpg", "", "gpg command to use")
	rootCmd.Flags().String("gpgflags", "", "Pass arguments to gpg")
	rootCmd.Flags().BoolP("help", "h", false, "Display syntax for the given operation")
	rootCmd.Flags().String("makepkg", "", "makepkg command to use")
	rootCmd.Flags().String("makepkgconf", "", "makepkg.conf file to use")
	rootCmd.Flags().String("mflags", "", "Pass arguments to makepkg")
	rootCmd.Flags().Bool("nomakepkgconf", false, "Use the default makepkg.conf")
	rootCmd.Flags().String("pacman", "", "pacman command to use")
	rootCmd.Flags().BoolP("repo", "N", false, "Assume targets are from the repositories")
	rootCmd.Flags().String("requestsplitn", "", "Max amount of packages to query per AUR request")
	rootCmd.Flags().Bool("save", false, "Causes the following options to be saved back to the config file when used")
	rootCmd.Flags().String("searchby", "", "Search for packages using a specified field")
	rootCmd.Flags().String("sortby", "", "Sort AUR results by a specific field during search")
	rootCmd.Flags().String("sudo", "", "sudo command to use")
	rootCmd.Flags().String("sudoflags", "", "Pass arguments to sudo")
	rootCmd.Flags().Bool("sudoloop", false, "Loop sudo calls in the background to avoid timeout")
	rootCmd.Flags().BoolP("version", "V", false, "show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"builddir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
		"editor":   carapace.ActionExecutables(),
		"git":      carapace.ActionExecutables(),
		"gpg":      carapace.ActionExecutables(),
		"makepkg":  carapace.ActionExecutables(),
		"pacman":   carapace.ActionExecutables(),
		"searchby": carapace.ActionValues("name-desc", "name", "maintainer", "submitter", "depends", "makedepends", "optdepends", "checkdepends", "provides", "conflicts", "replaces", "groups", "keywords", "comaintainers"),
		"sortby":   carapace.ActionValues("base", "modified", "name", "popularity", "submitted", "votes"),
		"sudo":     carapace.ActionExecutables(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		yay.ActionPackageSearch(),
	)

	embed.SubcommandsAsFlags(rootCmd,
		buildCmd,
		databaseCmd,
		deptestCmd,
		filesCmd,
		getpkgbuildCmd,
		queryCmd,
		removeCmd,
		showCmd,
		syncCmd,
		upgradeCmd,
		webCmd,
		yayCmd,
	)
}
