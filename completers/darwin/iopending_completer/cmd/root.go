package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iopending",
	Short: "print disk I/O pending events",
	Long:  "https://man.freebsd.org/cgi/man.cgi?iopending",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("d", "d", "", "Device instance")
	rootCmd.Flags().StringS("f", "f", "", "Filename")
	rootCmd.Flags().StringS("m", "m", "", "Mount point")
}
