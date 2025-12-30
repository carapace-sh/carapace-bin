package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-uhelper [OPTIONS] [MODE] [ARGUMENTS]",
	Short: "XBPS helper utilities",
	Long:  "https://man.voidlinux.org/xbps-uhelper",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().StringP("config", "C", "", "Path to confdir (xbps.d)")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Debug mode shown to stderr")
	rootCmd.PersistentFlags().StringP("rootdir", "r", "", "Full path to rootdir")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose messages")
	rootCmd.PersistentFlags().BoolP("version", "V", false, "Show XBPS version")
}
