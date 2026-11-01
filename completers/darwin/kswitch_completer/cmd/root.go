package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kswitch",
	Short: "switch Kerberos principal cache",
	Long:  "https://man.freebsd.org/cgi/man.cgi?kswitch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("c", "c", "", "Cache name")
	rootCmd.Flags().Bool("help", false, "Help")
	rootCmd.Flags().BoolS("i", "i", false, "Interactive")
	rootCmd.Flags().StringS("p", "p", "", "Principal")
	rootCmd.Flags().StringS("t", "t", "", "Type")
	rootCmd.Flags().Bool("version", false, "Version")
}
