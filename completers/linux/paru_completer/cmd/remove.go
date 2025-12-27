package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/paru_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pacman"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"R"},
	Short:   "Remove packages from the system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().String("arch", "", "set an alternate architecture")
	removeCmd.Flags().String("assume-installed", "", "add a virtual package to satisfy dependencies")
	removeCmd.Flags().String("cachedir", "", "set an alternate package cache location")
	removeCmd.Flags().BoolP("cascade", "c", false, "remove packages and all packages that depend on them")
	removeCmd.Flags().String("color", "", "colorize the output")
	removeCmd.Flags().String("config", "", "set an alternate configuration file")
	removeCmd.Flags().Bool("confirm", false, "always ask for confirmation")
	removeCmd.Flags().Bool("dbonly", false, "only modify database entries, not package files")
	removeCmd.Flags().StringP("dbpath", "b", "", "set an alternate database location")
	removeCmd.Flags().Bool("debug", false, "display debug messages")
	removeCmd.Flags().Bool("disable-download-timeout", false, "use relaxed timeouts for download")
	removeCmd.Flags().String("gpgdir", "", "set an alternate home directory for GnuPG")
	removeCmd.Flags().String("hookdir", "", "set an alternate hook location")
	removeCmd.Flags().String("logfile", "", "set an alternate log file")
	removeCmd.Flags().Bool("noconfirm", false, "do not ask for any confirmation")
	removeCmd.Flags().CountP("nodeps", "d", "skip dependency version checks (-dd to skip all checks)")
	removeCmd.Flags().Bool("noprogressbar", false, "do not show a progress bar when downloading files")
	removeCmd.Flags().BoolP("nosave", "n", false, "remove configuration files")
	removeCmd.Flags().Bool("noscriptlet", false, "do not execute the install scriptlet if one exists")
	removeCmd.Flags().BoolP("print", "p", false, "print the targets instead of performing the operation")
	removeCmd.Flags().String("print-format", "", "specify how the targets should be printed")
	removeCmd.Flags().CountP("recursive", "s", "remove unnecessary dependencies")
	removeCmd.Flags().StringP("root", "r", "", "set an alternate installation root")
	removeCmd.Flags().Bool("sysroot", false, "operate on a mounted guest system (root-only)")
	removeCmd.Flags().BoolP("unneeded", "u", false, "remove unneeded packages")
	removeCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	common.AddNewFlags(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"arch":     carapace.ActionValues("i686", "x86_64"),
		"cachedir": carapace.ActionDirectories(),
		"color":    carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"config":   carapace.ActionFiles(),
		"dbpath":   carapace.ActionFiles(),
		"gpgdir":   carapace.ActionDirectories(),
		"hookdir":  carapace.ActionDirectories(),
		"logfile":  carapace.ActionFiles(),
		"root":     carapace.ActionDirectories(),
	})

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		pacman.ActionPackages().FilterArgs(),
	)
}
