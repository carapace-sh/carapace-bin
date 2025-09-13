package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/paru_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pacman"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var databaseCmd = &cobra.Command{
	Use:     "database",
	Aliases: []string{"D"},
	Short:   "Operate on the package database",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databaseCmd).Standalone()

	databaseCmd.Flags().String("arch", "", "set an alternate architecture")
	databaseCmd.Flags().Bool("asdeps", false, "mark packages as non-explicitly installed")
	databaseCmd.Flags().Bool("asexplicit", false, "mark packages as explicitly installed")
	databaseCmd.Flags().String("cachedir", "", "set an alternate package cache location")
	databaseCmd.Flags().CountP("check", "k", "test local database for validity (-kk for sync databases)")
	databaseCmd.Flags().String("color", "", "colorize the output")
	databaseCmd.Flags().String("config", "", "set an alternate configuration file")
	databaseCmd.Flags().Bool("confirm", false, "always ask for confirmation")
	databaseCmd.Flags().StringP("dbpath", "b", "", "set an alternate database location")
	databaseCmd.Flags().Bool("debug", false, "display debug messages")
	databaseCmd.Flags().Bool("disable-download-timeout", false, "use relaxed timeouts for download")
	databaseCmd.Flags().String("gpgdir", "", "set an alternate home directory for GnuPG")
	databaseCmd.Flags().String("hookdir", "", "set an alternate hook location")
	databaseCmd.Flags().String("logfile", "", "set an alternate log file")
	databaseCmd.Flags().Bool("noconfirm", false, "do not ask for any confirmation")
	databaseCmd.Flags().BoolP("quiet", "q", false, "suppress output of success messages")
	databaseCmd.Flags().StringP("root", "r", "", "set an alternate installation root")
	databaseCmd.Flags().Bool("sysroot", false, "operate on a mounted guest system (root-only)")
	databaseCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	common.AddNewFlags(databaseCmd)

	carapace.Gen(databaseCmd).FlagCompletion(carapace.ActionMap{
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

	carapace.Gen(databaseCmd).PositionalAnyCompletion(
		pacman.ActionPackages().FilterArgs(),
	)
}
