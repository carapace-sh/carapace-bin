package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pacman"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:     "query",
	Aliases: []string{"Q"},
	Short:   "Query the package database",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queryCmd).Standalone()

	queryCmd.Flags().String("arch", "", "set an alternate architecture")
	queryCmd.Flags().String("cachedir", "", "set an alternate package cache location")
	queryCmd.Flags().BoolP("changelog", "c", false, "view the changelog of a package")
	queryCmd.Flags().CountP("check", "k", "check that package files exist (-kk for file properties)")
	queryCmd.Flags().String("color", "", "colorize the output")
	queryCmd.Flags().String("config", "", "set an alternate configuration file")
	queryCmd.Flags().Bool("confirm", false, "always ask for confirmation")
	queryCmd.Flags().StringP("dbpath", "b", "", "set an alternate database location")
	queryCmd.Flags().Bool("debug", false, "display debug messages")
	queryCmd.Flags().BoolP("deps", "d", false, "list packages installed as dependencies [filter]")
	queryCmd.Flags().Bool("disable-download-timeout", false, "use relaxed timeouts for download")
	queryCmd.Flags().Bool("disable-sandbox", false, "disables all sandbox features used for the downloader process")
	queryCmd.Flags().Bool("disable-sandbox-filesystem", false, "disables the filesystem part of the downloader process sandbox")
	queryCmd.Flags().Bool("disable-sandbox-syscalls", false, "disables the syscalls part of the downloader process sandbox")
	queryCmd.Flags().BoolP("explicit", "e", false, "list packages explicitly installed [filter]")
	queryCmd.Flags().StringP("file", "p", "", "query a package file instead of the database")
	queryCmd.Flags().BoolP("foreign", "m", false, "list installed packages not found in sync db(s) [filter]")
	queryCmd.Flags().String("gpgdir", "", "set an alternate home directory for GnuPG")
	queryCmd.Flags().BoolP("groups", "g", false, "view all members of a package group")
	queryCmd.Flags().String("hookdir", "", "set an alternate hook location")
	queryCmd.Flags().CountP("info", "i", "view package information (-ii for backup files)")
	queryCmd.Flags().BoolP("list", "l", false, "list the files owned by the queried package")
	queryCmd.Flags().String("logfile", "", "set an alternate log file")
	queryCmd.Flags().BoolP("native", "n", false, "list installed packages only found in sync db(s) [filter]")
	queryCmd.Flags().Bool("noconfirm", false, "do not ask for any confirmation")
	queryCmd.Flags().StringP("owns", "o", "", "query the package that owns <file>")
	queryCmd.Flags().BoolP("quiet", "q", false, "show less information for query and search")
	queryCmd.Flags().StringP("root", "r", "", "set an alternate installation root")
	queryCmd.Flags().StringP("search", "s", "", "search locally-installed packages for matching strings")
	queryCmd.Flags().Bool("sysroot", false, "operate on a mounted guest system (root-only)")
	queryCmd.Flags().CountP("unrequired", "t", "list packages not (optionally) required by any package")
	queryCmd.Flags().BoolP("upgrades", "u", false, "list outdated packages [filter]")
	queryCmd.Flags().BoolP("verbose", "v", false, "be verbose")

	carapace.Gen(queryCmd).FlagCompletion(carapace.ActionMap{
		"arch":     carapace.ActionValues("i686", "x86_64"),
		"cachedir": carapace.ActionDirectories(),
		"color":    carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"config":   carapace.ActionFiles(),
		"dbpath":   carapace.ActionFiles(),
		"file":     carapace.ActionFiles(),
		"gpgdir":   carapace.ActionDirectories(),
		"hookdir":  carapace.ActionDirectories(),
		"logfile":  carapace.ActionFiles(),
		"owns":     carapace.ActionFiles(),
		"root":     carapace.ActionDirectories(),
	})

	carapace.Gen(queryCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if queryCmd.Flag("file").Changed {
				return carapace.ActionFiles()
			}
			return pacman.ActionPackages().FilterArgs()
		}),
	)
}
