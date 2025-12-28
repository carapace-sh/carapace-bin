package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-digest [options] [file] [file+N]",
	Short: "XBPS utility to generate message digests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Show usage")
	rootCmd.Flags().StringP("mode", "m", "sha256", "Selects the digest mode, sha256 (default)")
	rootCmd.Flags().BoolP("version", "V", false, "Show XBPS version")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
