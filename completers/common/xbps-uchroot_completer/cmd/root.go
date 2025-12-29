package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-uchroot [OPTIONS] [--] <dir> <cmd> [<cmdargs>]",
	Short: "XBPS utility to chroot and bind mount with Linux namespaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("bind-ro", "B", "", "Bind mounts <src> into <dir>/<dest> (read-only)")
	rootCmd.Flags().StringP("bind-rw", "b", "", "Bind mounts <src> into <dir>/<dest> (read-write)")
	rootCmd.Flags().BoolP("overlayfs", "O", false, "Creates a tempdir and mounts <dir> read-only via overlayfs")
	rootCmd.Flags().BoolP("tmpfs", "t", false, "Creates a tempdir and mounts <dir> on tmpfs (for use with -O)")
	rootCmd.Flags().StringP("options", "o", "", "Options to be passed to the tmpfs mount (for use with -t)")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().BoolP("verbose", "v", false, "Show XBPS version")

	carapace.Gen(rootCmd).PositionalCompletion(carapace.ActionDirectories())
}
