package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hostname",
	Short: "set or print name of current host system",
	Long:  "https://keith.github.io/xcode-manpages/hostname.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("d", "d", false, "Show DNS domain name")
	rootCmd.Flags().BoolS("f", "f", false, "Show FQDN")
	rootCmd.Flags().BoolS("s", "s", false, "Show short hostname")
}
