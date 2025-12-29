package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-uunshare [OPTIONS] [--] <dir> <cmd> [<cmdargs>]",
	Short: "XBPS utility to chroot and bind mount with Linux user namespaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("bind-rw", "b", "", "Bind mounts <src> into <dir>/<dest> (read-write)")
	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().BoolP("verbose", "v", false, "Show XBPS version")

	carapace.Gen(rootCmd).PositionalCompletion(carapace.ActionDirectories())
}
