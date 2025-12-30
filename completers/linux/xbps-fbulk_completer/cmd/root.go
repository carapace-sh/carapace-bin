package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xbps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-fbulk [OPTIONS] /path/to/void-packages [pkg pkg+N]",
	Short: "XBPS utility to perform a fast bulk build of void-packages",
	Long:  "https://man.voidlinux.org/xbps-fbulk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().IntP("jobs", "j", 0, "Number of parallel builds")
	rootCmd.Flags().StringP("logdir", "l", "", "Path to store logs")
	rootCmd.Flags().BoolP("system", "s", false, "System rebuild mode")
	rootCmd.Flags().BoolP("verbose", "V", false, "Enable verbose mode")
	rootCmd.Flags().BoolP("version", "v", false, "Show XBPS version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"logdir": carapace.ActionDirectories(),
	})
	carapace.Gen(rootCmd).PositionalCompletion(carapace.ActionDirectories())
	carapace.Gen(rootCmd).PositionalAnyCompletion(xbps.ActionPackages().FilterArgs())
}
