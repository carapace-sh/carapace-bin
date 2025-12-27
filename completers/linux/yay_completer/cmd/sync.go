package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pacman"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yay"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:     "sync",
	Aliases: []string{"S"},
	Short:   "Synchronize packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncCmd).Standalone()

	syncCmd.Flags().String("arch", "", "set an alternate architecture")
	syncCmd.Flags().Bool("asdeps", false, "install packages as non-explicitly installed")
	syncCmd.Flags().Bool("asexplicit", false, "install packages as explicitly installed")
	syncCmd.Flags().String("assume-installed", "", "add a virtual package to satisfy dependencies")
	syncCmd.Flags().String("cachedir", "", "set an alternate package cache location")
	syncCmd.Flags().CountP("clean", "c", "remove old packages from cache directory (-cc for all)")
	syncCmd.Flags().String("color", "", "colorize the output")
	syncCmd.Flags().String("config", "", "set an alternate configuration file")
	syncCmd.Flags().Bool("confirm", false, "always ask for confirmation")
	syncCmd.Flags().Bool("dbonly", false, "only modify database entries, not package files")
	syncCmd.Flags().StringP("dbpath", "b", "", "set an alternate database location")
	syncCmd.Flags().Bool("debug", false, "display debug messages")
	syncCmd.Flags().Bool("disable-download-timeout", false, "use relaxed timeouts for download")
	syncCmd.Flags().BoolP("downloadonly", "w", false, "download packages but do not install/upgrade anything")
	syncCmd.Flags().String("gpgdir", "", "set an alternate home directory for GnuPG")
	syncCmd.Flags().CountP("groups", "g", "view all members of a package group")
	syncCmd.Flags().String("hookdir", "", "set an alternate hook location")
	syncCmd.Flags().StringSlice("ignore", nil, "ignore a package upgrade")
	syncCmd.Flags().StringSlice("ignoregroup", nil, "ignore a group upgrade")
	syncCmd.Flags().BoolP("info", "i", false, "view package information (-ii for extended information)")
	syncCmd.Flags().StringP("list", "l", "", "view a list of packages in a repo")
	syncCmd.Flags().String("logfile", "", "set an alternate log file")
	syncCmd.Flags().Bool("needed", false, "do not reinstall up to date packages")
	syncCmd.Flags().Bool("noconfirm", false, "do not ask for any confirmation")
	syncCmd.Flags().BoolP("nodeps", "d", false, "skip dependency version checks (-dd to skip all checks)")
	syncCmd.Flags().Bool("noprogressbar", false, "do not show a progress bar when downloading files")
	syncCmd.Flags().Bool("noscriptlet", false, "do not execute the install scriptlet if one exists")
	syncCmd.Flags().StringSlice("overwrite", nil, "overwrite conflicting files")
	syncCmd.Flags().BoolP("print", "p", false, "print the targets instead of performing the operation")
	syncCmd.Flags().String("print-format", "", "specify how the targets should be printed")
	syncCmd.Flags().BoolP("quiet", "q", false, "show less information for query and search")
	syncCmd.Flags().CountP("refresh", "y", "download fresh package databases from the server")
	syncCmd.Flags().StringP("root", "r", "", "set an alternate installation root")
	syncCmd.Flags().StringP("search", "s", "", "search remote repositories for matching strings")
	syncCmd.Flags().Bool("sysroot", false, "operate on a mounted guest system (root-only)")
	syncCmd.Flags().CountP("sysupgrade", "u", "upgrade installed packages (-uu enables downgrades)")
	syncCmd.Flags().BoolP("verbose", "v", false, "be verbose")

	carapace.Gen(syncCmd).FlagCompletion(carapace.ActionMap{
		"arch":     carapace.ActionValues("i686", "x86_64"),
		"cachedir": carapace.ActionDirectories(),
		"color":    carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"config":   carapace.ActionFiles(),
		"dbpath":   carapace.ActionFiles(),
		"gpgdir":   carapace.ActionDirectories(),
		"hookdir":  carapace.ActionDirectories(),
		// TODO list
		"ignore":      yay.ActionPackageSearch().UniqueList(","),
		"ignoregroup": pacman.ActionPackageGroups().UniqueList(","),
		"logfile":     carapace.ActionFiles(),
		"overwrite": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return pacman.ActionPackageFiles(c.Args...).List(",")
			}
			return carapace.ActionValues()
		}),
		"root": carapace.ActionDirectories(),
	})

	carapace.Gen(syncCmd).PositionalAnyCompletion(
		yay.ActionPackageSearch().FilterArgs(),
	)
}
