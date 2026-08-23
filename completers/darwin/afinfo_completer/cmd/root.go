package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "afinfo",
	Short: "Audio File Info",
	Long:  "https://man.freebsd.org/cgi/man.cgi?afinfo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("b", "b", false, "Brief info")
	rootCmd.Flags().BoolS("h", "h", false, "Print help text")
	rootCmd.Flags().BoolS("l", "l", false, "List info")
	rootCmd.Flags().BoolS("r", "r", false, "Real time")
}
