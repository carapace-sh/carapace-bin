package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:     "files",
	Aliases: []string{"F"},
	Short:   "Query the files database",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(filesCmd).Standalone()

	filesCmd.Flags().String("arch", "", "set an alternate architecture")
	filesCmd.Flags().String("cachedir", "", "set an alternate package cache location")
	filesCmd.Flags().String("color", "", "colorize the output")
	filesCmd.Flags().String("config", "", "set an alternate configuration file")
	filesCmd.Flags().Bool("confirm", false, "always ask for confirmation")
	filesCmd.Flags().StringP("dbpath", "b", "", "set an alternate database location")
	filesCmd.Flags().Bool("debug", false, "display debug messages")
	filesCmd.Flags().Bool("disable-download-timeout", false, "use relaxed timeouts for download")
	filesCmd.Flags().Bool("disable-sandbox", false, "disables all sandbox features used for the downloader process")
	filesCmd.Flags().Bool("disable-sandbox-filesystem", false, "disables the filesystem part of the downloader process sandbox")
	filesCmd.Flags().Bool("disable-sandbox-syscalls", false, "disables the syscalls part of the downloader process sandbox")
	filesCmd.Flags().String("gpgdir", "", "set an alternate home directory for GnuPG")
	filesCmd.Flags().String("hookdir", "", "set an alternate hook location")
	filesCmd.Flags().BoolP("list", "l", false, "list the files owned by the queried package")
	filesCmd.Flags().String("logfile", "", "set an alternate log file")
	filesCmd.Flags().Bool("machinereadable", false, "produce machine-readable output")
	filesCmd.Flags().Bool("noconfirm", false, "do not ask for any confirmation")
	filesCmd.Flags().BoolP("quiet", "q", false, "show less information for query and search")
	filesCmd.Flags().CountP("refresh", "y", "download fresh package databases from the server")
	filesCmd.Flags().BoolP("regex", "x", false, "enable searching using regular expressions")
	filesCmd.Flags().StringP("root", "r", "", "set an alternate installation root")
	filesCmd.Flags().Bool("sysroot", false, "operate on a mounted guest system (root-only)")
	filesCmd.Flags().BoolP("verbose", "v", false, "be verbose")

	carapace.Gen(filesCmd).FlagCompletion(carapace.ActionMap{
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

	carapace.Gen(filesCmd).PositionalAnyCompletion(
		carapace.ActionFiles().Chdir("/"),
	)
}
