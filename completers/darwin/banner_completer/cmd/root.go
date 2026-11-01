package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "banner",
	Short: "print large banner on printer",
	Long:  "https://man.freebsd.org/cgi/man.cgi?banner",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("d", "d", false, "Enable debug")
	rootCmd.Flags().BoolS("t", "t", false, "Enable trace")
	rootCmd.Flags().StringS("w", "w", "", "Width")
}
