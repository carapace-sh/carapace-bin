package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pacman"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"U"},
	Short:   "Upgrade or add packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().String("arch", "", "set an alternate architecture")
	upgradeCmd.Flags().Bool("asdeps", false, "install packages as non-explicitly installed")
	upgradeCmd.Flags().Bool("asexplicit", false, "install packages as explicitly installed")
	upgradeCmd.Flags().String("assume-installed", "", "add a virtual package to satisfy dependencies")
	upgradeCmd.Flags().String("cachedir", "", "set an alternate package cache location")
	upgradeCmd.Flags().String("color", "", "colorize the output")
	upgradeCmd.Flags().String("config", "", "set an alternate configuration file")
	upgradeCmd.Flags().Bool("confirm", false, "always ask for confirmation")
	upgradeCmd.Flags().Bool("dbonly", false, "only modify database entries, not package files")
	upgradeCmd.Flags().StringP("dbpath", "b", "", "set an alternate database location")
	upgradeCmd.Flags().Bool("debug", false, "display debug messages")
	upgradeCmd.Flags().Bool("disable-download-timeout", false, "use relaxed timeouts for download")
	upgradeCmd.Flags().BoolP("downloadonly", "w", false, "download packages but do not install/upgrade anything")
	upgradeCmd.Flags().String("gpgdir", "", "set an alternate home directory for GnuPG")
	upgradeCmd.Flags().String("hookdir", "", "set an alternate hook location")
	upgradeCmd.Flags().StringSlice("ignore", nil, "ignore a package upgrade")
	upgradeCmd.Flags().StringSlice("ignoregroup", nil, "ignore a group upgrade")
	upgradeCmd.Flags().String("logfile", "", "set an alternate log file")
	upgradeCmd.Flags().Bool("needed", false, "do not reinstall up to date packages")
	upgradeCmd.Flags().Bool("noconfirm", false, "do not ask for any confirmation")
	upgradeCmd.Flags().CountP("nodeps", "d", "skip dependency version checks (-dd to skip all checks)")
	upgradeCmd.Flags().Bool("noprogressbar", false, "do not show a progress bar when downloading files")
	upgradeCmd.Flags().Bool("noscriptlet", false, "do not execute the install scriptlet if one exists")
	upgradeCmd.Flags().StringSlice("overwrite", nil, "overwrite conflicting files")
	upgradeCmd.Flags().BoolP("print", "p", false, "print the targets instead of performing the operation")
	upgradeCmd.Flags().String("print-format", "", "specify how the targets should be printed")
	upgradeCmd.Flags().StringP("root", "r", "", "set an alternate installation root")
	upgradeCmd.Flags().Bool("sysroot", false, "operate on a mounted guest system (root-only)")
	upgradeCmd.Flags().BoolP("verbose", "v", false, "be verbose")

	carapace.Gen(upgradeCmd).FlagCompletion(carapace.ActionMap{
		"arch":        carapace.ActionValues("i686", "x86_64"),
		"cachedir":    carapace.ActionDirectories(),
		"color":       carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"config":      carapace.ActionFiles(),
		"dbpath":      carapace.ActionFiles(),
		"gpgdir":      carapace.ActionDirectories(),
		"hookdir":     carapace.ActionDirectories(),
		"ignore":      pacman.ActionPackageSearch().UniqueList(","),
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

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
