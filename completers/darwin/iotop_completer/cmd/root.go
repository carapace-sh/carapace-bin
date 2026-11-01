package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iotop",
	Short: "display top disk I/O events by process",
	Long:  "https://man.freebsd.org/cgi/man.cgi?iotop",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("C", "C", false, "Don't clear screen")
	rootCmd.Flags().BoolS("D", "D", false, "Print delta times")
	rootCmd.Flags().BoolS("P", "P", false, "Print percent I/O")
	rootCmd.Flags().BoolS("Z", "Z", false, "Print zone ID")
	rootCmd.Flags().StringS("d", "d", "", "Device instance")
	rootCmd.Flags().StringS("f", "f", "", "Filename")
	rootCmd.Flags().BoolS("j", "j", false, "Print project ID")
	rootCmd.Flags().StringS("m", "m", "", "Mount point")
	rootCmd.Flags().BoolS("o", "o", false, "Print disk delta times")
	rootCmd.Flags().StringS("t", "t", "", "Top number only")
}
