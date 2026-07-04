package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yay"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var yayCmd = &cobra.Command{
	Use:     "yay",
	Aliases: []string{"Y"},
	Short:   "YAY specific operations",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(yayCmd).Standalone()

	yayCmd.Flags().String("answerclean", "", "Set a predetermined answer for the clean build menu")
	yayCmd.Flags().String("answerdiff", "", "Set a predetermined answer for the diff menu")
	yayCmd.Flags().String("answeredit", "", "Set a predetermined answer for the edit pkgbuild menu")
	yayCmd.Flags().String("answerupgrade", "", "Set a predetermined answer for the upgrade menu")
	yayCmd.Flags().Bool("arch", false, "Assume targets are from the repositories")
	yayCmd.Flags().Bool("asdeps", false, "Install packages as non-explicitly installed")
	yayCmd.Flags().Bool("asexplicit", false, "Install packages as explicitly installed")
	yayCmd.Flags().Bool("askremovemake", false, "Ask to remove makedepends after install")
	yayCmd.Flags().Bool("askyesremovemake", false, "Ask to remove makedepends after install (Y as default)")
	yayCmd.Flags().String("assume-installed", "", "Add a virtual package to satisfy dependencies")
	yayCmd.Flags().String("aurrpcurl", "", "Set an alternative URL for the AUR /rpc endpoint")
	yayCmd.Flags().String("aururl", "", "Set an alternative AUR URL")
	yayCmd.Flags().Bool("bottomup", false, "Shows AUR's packages first and then repository's")
	yayCmd.Flags().String("builddir", "", "Directory used to download and run PKGBUILDS")
	yayCmd.Flags().String("cachedir", "", "Directory used to store downloaded packages")
	yayCmd.Flags().BoolP("clean", "c", false, "Remove unneeded dependencies")
	yayCmd.Flags().Bool("cleanafter", false, "Remove package sources after successful install")
	yayCmd.Flags().Bool("cleanmenu", false, "Give the option to clean build PKGBUILDS")
	yayCmd.Flags().String("color", "", "Colorize the output")
	yayCmd.Flags().String("completioninterval", "", "Time in days to refresh completion cache")
	yayCmd.Flags().String("config", "", "pacman.conf file to use")
	yayCmd.Flags().Bool("confirm", false, "Always prompt for confirmation")
	yayCmd.Flags().String("dbonly", "", "Only modify database entries")
	yayCmd.Flags().Bool("debug", false, "Display debug messages")
	yayCmd.Flags().Bool("devel", false, "Check development packages during sysupgrade")
	yayCmd.Flags().Bool("diffmenu", false, "Give the option to show diffs for build files")
	yayCmd.Flags().Bool("disable-download-timeout", false, "Disable download timeout")
	yayCmd.Flags().Bool("disable-sandbox", false, "Disable sandboxing")
	yayCmd.Flags().Bool("disable-sandbox-filesystem", false, "Disable sandbox filesystem")
	yayCmd.Flags().Bool("disable-sandbox-syscalls", false, "Disable sandbox syscalls")
	yayCmd.Flags().Bool("doublelineresults", false, "List each search result on two lines, like pacman")
	yayCmd.Flags().Bool("editmenu", false, "Give the option to edit/view PKGBUILDS")
	yayCmd.Flags().String("editor", "", "Editor to use when editing PKGBUILDs")
	yayCmd.Flags().String("editorflags", "", "Pass arguments to editor")
	yayCmd.Flags().Bool("gendb", false, "Generates development package DB used for updating")
	yayCmd.Flags().String("git", "", "git command to use")
	yayCmd.Flags().String("gitflags", "", "Pass arguments to git")
	yayCmd.Flags().String("gpg", "", "gpg command to use")
	yayCmd.Flags().String("gpgdir", "", "GPG directory to use")
	yayCmd.Flags().String("gpgflags", "", "Pass arguments to gpg")
	yayCmd.Flags().String("hookdir", "", "Hook directory to use")
	yayCmd.Flags().String("ignore", "", "Ignore a package upgrade")
	yayCmd.Flags().String("ignoregroup", "", "Ignore a group upgrade")
	yayCmd.Flags().Bool("keepsrc", false, "Keep pkg/ and src/ after building packages")
	yayCmd.Flags().String("logfile", "", "Specify log file")
	yayCmd.Flags().String("makepkg", "", "makepkg command to use")
	yayCmd.Flags().String("makepkgconf", "", "makepkg.conf file to use")
	yayCmd.Flags().String("mflags", "", "Pass arguments to makepkg")
	yayCmd.Flags().Bool("needed", false, "Do not reinstall the targets already up to date")
	yayCmd.Flags().Bool("noanswerclean", false, "Unset the answer for the clean build menu")
	yayCmd.Flags().Bool("noanswerdiff", false, "Unset the answer for the edit diff menu")
	yayCmd.Flags().Bool("noansweredit", false, "Unset the answer for the edit pkgbuild menu")
	yayCmd.Flags().Bool("noanswerupgrade", false, "Unset the answer for the upgrade menu")
	yayCmd.Flags().Bool("noconfirm", false, "Do not ask for any confirmation")
	yayCmd.Flags().Bool("nomakepkgconf", false, "Use the default makepkg.conf")
	yayCmd.Flags().Bool("noprogressbar", false, "Do not show a progress bar when downloading")
	yayCmd.Flags().Bool("norebuild", false, "Skip package build if in cache and up to date")
	yayCmd.Flags().Bool("noredownload", false, "Skip pkgbuild download if in cache and up to date")
	yayCmd.Flags().Bool("noremovemake", false, "Don't remove makedepends after install")
	yayCmd.Flags().Bool("noscriptlet", false, "Do not try to install scriptlet")
	yayCmd.Flags().Bool("overwrite", false, "Overwrite conflicting files")
	yayCmd.Flags().String("pacman", "", "pacman command to use")
	yayCmd.Flags().Bool("pgpfetch", false, "Prompt to import PGP keys from PKGBUILDs")
	yayCmd.Flags().String("print-format", "", "Specify how the targets should be printed")
	yayCmd.Flags().Bool("provides", false, "Look for matching providers when searching for packages")
	yayCmd.Flags().Bool("rebuild", false, "Always build target packages")
	yayCmd.Flags().Bool("rebuildall", false, "Always build all AUR packages")
	yayCmd.Flags().Bool("rebuildtree", false, "Always build all AUR packages even if installed")
	yayCmd.Flags().Bool("redownload", false, "Always download pkgbuilds of targets")
	yayCmd.Flags().Bool("redownloadall", false, "Always download pkgbuilds of all AUR packages")
	yayCmd.Flags().Bool("removemake", false, "Remove makedepends after install")
	yayCmd.Flags().String("requestsplitn", "", "Max amount of packages to query per AUR request")
	yayCmd.Flags().Bool("save", false, "Causes the following options to be saved back to the config file")
	yayCmd.Flags().String("searchby", "", "Search for packages using a specified field")
	yayCmd.Flags().Bool("singlelineresults", false, "List each search result on its own line")
	yayCmd.Flags().String("sortby", "", "Sort AUR results by a specific field during search")
	yayCmd.Flags().String("sudo", "", "sudo command to use")
	yayCmd.Flags().String("sudoflags", "", "Pass arguments to sudo")
	yayCmd.Flags().Bool("sudoloop", false, "Loop sudo calls in the background to avoid timeout")
	yayCmd.Flags().String("sysroot", "", "Specify an alternative system root")
	yayCmd.Flags().Bool("topdown", false, "Shows repository's packages first and then AUR's")
	yayCmd.Flags().Bool("useask", false, "Automatically resolve conflicts using pacman's ask flag")

	carapace.Gen(yayCmd).FlagCompletion(carapace.ActionMap{
		"answerclean":   carapace.ActionValues("Yes", "No", "All", "None", "Editor", "Diff"),
		"answerdiff":    carapace.ActionValues("Yes", "No", "All", "None", "Editor", "Diff"),
		"answeredit":    carapace.ActionValues("Yes", "No", "All", "None", "Editor", "Diff"),
		"answerupgrade": carapace.ActionValues("Yes", "No", "All", "None", "Editor", "Diff"),
		"builddir":      carapace.ActionDirectories(),
		"cachedir":      carapace.ActionDirectories(),
		"color":         carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"config":        carapace.ActionFiles(),
		"editor":        carapace.ActionExecutables(),
		"git":           carapace.ActionExecutables(),
		"gpg":           carapace.ActionExecutables(),
		"gpgdir":        carapace.ActionDirectories(),
		"hookdir":       carapace.ActionDirectories(),
		"logfile":       carapace.ActionFiles(),
		"makepkg":       carapace.ActionExecutables(),
		"makepkgconf":   carapace.ActionFiles(),
		"pacman":        carapace.ActionExecutables(),
		"searchby":      carapace.ActionValues("name-desc", "name", "maintainer", "submitter", "depends", "makedepends", "optdepends", "checkdepends", "provides", "conflicts", "replaces", "groups", "keywords", "comaintainers"),
		"sortby":        carapace.ActionValues("base", "modified", "name", "popularity", "submitted", "votes"),
		"sudo":          carapace.ActionExecutables(),
		"sysroot":       carapace.ActionDirectories(),
	})

	carapace.Gen(yayCmd).PositionalAnyCompletion(
		yay.ActionPackageSearch().FilterArgs(),
	)
}
