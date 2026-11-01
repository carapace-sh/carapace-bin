package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iosnoop",
	Short: "print disk I/O events as they occur",
	Long:  "https://man.freebsd.org/cgi/man.cgi?iosnoop",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "Dump all data, space delimited")
	rootCmd.Flags().BoolS("D", "D", false, "Print time delta")
	rootCmd.Flags().BoolS("N", "N", false, "Print major and minor numbers")
	rootCmd.Flags().BoolS("a", "a", false, "Print all data")
	rootCmd.Flags().StringS("d", "d", "", "Device instance")
	rootCmd.Flags().BoolS("e", "e", false, "Print device name")
	rootCmd.Flags().StringS("f", "f", "", "Filename")
	rootCmd.Flags().BoolS("i", "i", false, "Print device instance")
	rootCmd.Flags().StringS("m", "m", "", "Mount point")
	rootCmd.Flags().StringS("n", "n", "", "Name")
	rootCmd.Flags().BoolS("o", "o", false, "Print disk delta time")
	rootCmd.Flags().BoolS("s", "s", false, "Print start time")
	rootCmd.Flags().BoolS("t", "t", false, "Print completion time")
	rootCmd.Flags().BoolS("v", "v", false, "Print completion time string")
}
