package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pacman"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var deptestCmd = &cobra.Command{
	Use:     "deptest",
	Aliases: []string{"T"},
	Short:   "Check dependencies",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deptestCmd).Standalone()

	deptestCmd.Flags().String("arch", "", "set an alternate architecture")
	deptestCmd.Flags().String("cachedir", "", "set an alternate package cache location")
	deptestCmd.Flags().String("color", "", "colorize the output")
	deptestCmd.Flags().String("config", "", "set an alternate configuration file")
	deptestCmd.Flags().Bool("confirm", false, "always ask for confirmation")
	deptestCmd.Flags().StringP("dbpath", "b", "", "set an alternate database location")
	deptestCmd.Flags().Bool("debug", false, "display debug messages")
	deptestCmd.Flags().Bool("disable-download-timeout", false, "use relaxed timeouts for download")
	deptestCmd.Flags().String("gpgdir", "", "set an alternate home directory for GnuPG")
	deptestCmd.Flags().String("hookdir", "", "set an alternate hook location")
	deptestCmd.Flags().String("logfile", "", "set an alternate log file")
	deptestCmd.Flags().Bool("noconfirm", false, "do not ask for any confirmation")
	deptestCmd.Flags().StringP("root", "r", "", "set an alternate installation root")
	deptestCmd.Flags().Bool("sysroot", false, "operate on a mounted guest system (root-only)")
	deptestCmd.Flags().BoolP("verbose", "v", false, "be verbose")

	carapace.Gen(deptestCmd).FlagCompletion(carapace.ActionMap{
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

	carapace.Gen(deptestCmd).PositionalAnyCompletion(
		pacman.ActionPackages().FilterArgs(),
	)
}
