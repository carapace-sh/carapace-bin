package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/paru_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/paru"
	"github.com/carapace-sh/carapace-bin/pkg/util/embed"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "paru",
	Short: "Feature packed AUR helper",
	Long:  "https://github.com/Morganamilo/paru",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("arch", "", "set an alternate architecture")
	rootCmd.Flags().String("ask", "", "pre-specify answers for questions")
	rootCmd.Flags().String("cachedir", "", "set an alternate package cache location")
	rootCmd.Flags().CountP("clean", "c", "Remove unneeded dependencies")
	rootCmd.Flags().String("color", "", "colorize the output")
	rootCmd.Flags().String("config", "", "set an alternate configuration file")
	rootCmd.Flags().Bool("confirm", false, "always ask for confirmation")
	rootCmd.Flags().StringP("dbpath", "b", "", "set an alternate database location")
	rootCmd.Flags().Bool("debug", false, "display debug messages")
	rootCmd.Flags().Bool("disable-download-timeout", false, "use relaxed timeouts for download")
	rootCmd.Flags().Bool("gendb", false, "Generates development package DB used for updating")
	rootCmd.Flags().String("gpgdir", "", "set an alternate home directory for GnuPG")
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().String("hookdir", "", "set an alternate hook location")
	rootCmd.Flags().String("logfile", "", "set an alternate log file")
	rootCmd.Flags().Bool("noconfirm", false, "do not ask for any confirmation")
	rootCmd.Flags().StringP("root", "r", "", "set an alternate installation root")
	rootCmd.Flags().Bool("sysroot", false, "operate on a mounted guest system (root-only)")
	rootCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	rootCmd.Flags().BoolP("version", "V", false, "show version")
	common.AddNewFlags(rootCmd)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
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

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			a := paru.ActionPackageSearch()
			rootCmd.Flags().Visit(func(f *pflag.Flag) {
				if f.Changed {
					a = carapace.ActionValues()
				}
			})
			return a
		}),
	)

	embed.SubcommandsAsFlags(rootCmd,
		buildCmd,
		chrootctlCmd,
		databaseCmd,
		deptestCmd,
		filesCmd,
		getpkgbuildCmd,
		queryCmd,
		removeCmd,
		repoctlCmd,
		showCmd,
		syncCmd,
		upgradeCmd,
	)
}
