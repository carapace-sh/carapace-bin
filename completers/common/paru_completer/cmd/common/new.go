package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddNewFlags(cmd *cobra.Command) {
	cmd.Flags().String("asp", "", "asp command to use")
	cmd.Flags().BoolP("aur", "a", false, "Assume targets are from the AUR")
	cmd.Flags().String("aururl", "", "Set an alternative AUR URL")
	cmd.Flags().String("bat", "", "bat command to use")
	cmd.Flags().String("batflags", "", "Pass arguments to bat")
	cmd.Flags().String("clonedir", "", "Directory used to download and run PKGBUILDs")
	cmd.Flags().String("completioninterval", "", "Time in days to refresh completion cache")
	cmd.Flags().String("fm", "", "File manager to use for PKGBUILD review")
	cmd.Flags().String("fmflags", "", "Pass arguments to file manager")
	cmd.Flags().String("git", "", "git command to use")
	cmd.Flags().String("gitflags", "", "Pass arguments to git")
	cmd.Flags().String("gpg", "", "gpg command to use")
	cmd.Flags().String("gpgflags", "", "Pass arguments to gpg")
	cmd.Flags().String("limit", "", "Limits the number of items returned in a search")
	cmd.Flags().String("makepkg", "", "makepkg command to use")
	cmd.Flags().String("mflags", "", "Pass arguments to makepkg")
	cmd.Flags().String("pacman", "", "pacman command to use")
	cmd.Flags().BoolP("regex", "x", false, "Enable regex for aur search")
	cmd.Flags().Bool("repo", false, "Assume targets are from the repositories")
	cmd.Flags().Bool("review", false, "Don't skip the review process")
	cmd.Flags().String("searchby", "", "Search for packages using a specified field")
	cmd.Flags().Bool("skipreview", false, "Skip the review process")
	cmd.Flags().String("sortby", "", "Sort AUR results by a specific field during search")
	cmd.Flags().String("sudo", "", "sudo command to use")
	cmd.Flags().String("sudoflags", "", "Pass arguments to sudo")

	cmd.Flags().Bool("batchinstall", false, "Build multiple AUR packages then install them together")
	cmd.Flags().Bool("chroot", false, "Build packages in a chroot")
	cmd.Flags().Bool("cleanafter", false, "Remove package sources after install")
	cmd.Flags().Bool("combinedupgrade", false, "Refresh then perform the repo and AUR upgrade together")
	cmd.Flags().Bool("devel", false, "Check development packages during sysupgrade")
	cmd.Flags().Bool("failfast", false, "Exit as soon as building an AUR package fails")
	cmd.Flags().Bool("installdebug", false, "Also install debug packages when a package provides them")
	cmd.Flags().Bool("keepsrc", false, "Keep src/ and pkg/ dirs after building packages")
	cmd.Flags().Bool("newsonupgrade", false, "Print new news during sysupgrade")
	cmd.Flags().Bool("pgpfetch", false, "Prompt to import PGP keys from PKGBUILDs")
	cmd.Flags().Bool("provides", false, "Look for matching providers when searching for packages")
	cmd.Flags().Bool("rebuild", false, "Always build target packages")
	cmd.Flags().Bool("redownload", false, "Always download PKGBUILDs of targets")
	cmd.Flags().Bool("removemake", false, "Remove makedepends after install")
	cmd.Flags().Bool("savechanges", false, "Commit changes to pkgbuilds made during review")
	cmd.Flags().Bool("sign", false, "Sign packages with gpg")
	cmd.Flags().Bool("signdb", false, "Sign databases with gpg")
	cmd.Flags().Bool("sudoloop", false, "Loop sudo calls in the background to avoid timeout")
	cmd.Flags().Bool("upgrademenu", false, "Show interactive menu to skip upgrades")
	cmd.Flags().Bool("useask", false, "Automatically resolve conflicts using pacman's ask flag")

	cmd.Flags().Bool("nobatchinstall", false, "Build multiple AUR packages then install them together")
	cmd.Flags().Bool("nochroot", false, "Build packages in a chroot")
	cmd.Flags().Bool("nocleanafter", false, "Remove package sources after install")
	cmd.Flags().Bool("nocombinedupgrade", false, "Refresh then perform the repo and AUR upgrade together")
	cmd.Flags().Bool("nodevel", false, "Check development packages during sysupgrade")
	cmd.Flags().Bool("nofailfast", false, "Exit as soon as building an AUR package fails")
	cmd.Flags().Bool("noinstalldebug", false, "Also install debug packages when a package provides them")
	cmd.Flags().Bool("nokeepsrc", false, "Keep src/ and pkg/ dirs after building packages")
	cmd.Flags().Bool("nonewsonupgrade", false, "Print new news during sysupgrade")
	cmd.Flags().Bool("nopgpfetch", false, "Prompt to import PGP keys from PKGBUILDs")
	cmd.Flags().Bool("noprovides", false, "Look for matching providers when searching for packages")
	cmd.Flags().Bool("norebuild", false, "Always build target packages")
	cmd.Flags().Bool("noredownload", false, "Always download PKGBUILDs of targets")
	cmd.Flags().Bool("noremovemake", false, "Remove makedepends after install")
	cmd.Flags().Bool("nosavechanges", false, "Commit changes to pkgbuilds made during review")
	cmd.Flags().Bool("nosign", false, "Sign packages with gpg")
	cmd.Flags().Bool("nosigndb", false, "Sign databases with gpg")
	cmd.Flags().Bool("nosudoloop", false, "Loop sudo calls in the background to avoid timeout")
	cmd.Flags().Bool("noupgrademenu", false, "Show interactive menu to skip upgrades")
	cmd.Flags().Bool("nouseask", false, "Automatically resolve conflicts using pacman's ask flag")

	cmd.Flags().Bool("bottomup", false, "Shows AUR's packages first and then repository's")
	cmd.Flags().Bool("develsuffixes", false, "Suffixes used to decide if a package is a devel package")
	cmd.Flags().Bool("localrepo", false, "Build packages into a local repo")
	cmd.Flags().Bool("nocheck", false, "Don't resolve checkdepends or run the check function")
	cmd.Flags().Bool("topdown", false, "Shows repository's packages first and then AUR's")

	cmd.Flag("noupgrademenu").Hidden = true
	cmd.Flag("noremovemake").Hidden = true
	cmd.Flag("nocleanafter").Hidden = true
	cmd.Flag("norebuild").Hidden = true
	cmd.Flag("noredownload").Hidden = true
	cmd.Flag("nopgpfetch").Hidden = true
	cmd.Flag("nouseask").Hidden = true
	cmd.Flag("nosavechanges").Hidden = true
	cmd.Flag("nonewsonupgrade").Hidden = true
	cmd.Flag("nocombinedupgrade").Hidden = true
	cmd.Flag("nobatchinstall").Hidden = true
	cmd.Flag("noprovides").Hidden = true
	cmd.Flag("nodevel").Hidden = true
	cmd.Flag("noinstalldebug").Hidden = true
	cmd.Flag("nosudoloop").Hidden = true
	cmd.Flag("nochroot").Hidden = true
	cmd.Flag("nofailfast").Hidden = true
	cmd.Flag("nokeepsrc").Hidden = true
	cmd.Flag("nosign").Hidden = true
	cmd.Flag("nosigndb").Hidden = true

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"asp": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"bat": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"clonedir": carapace.ActionDirectories(),
		"fm": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"git": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"gpg": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"makepkg": carapace.ActionFiles(),
		"pacman": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"sudo": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		// TODO sortby, searchby
	})
}
