package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tftp",
	Short: "trivial file transfer program",
	Long:  "https://man.freebsd.org/cgi/man.cgi?tftp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("e", "e", false, "Set blocksize to largest supported value")
}
