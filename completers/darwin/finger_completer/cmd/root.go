package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "finger",
	Short: "user information lookup program",
	Long:  "https://man.freebsd.org/cgi/man.cgi?finger",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("4", "4", false, "Use IPv4 addresses only")
	rootCmd.Flags().BoolS("6", "6", false, "Use IPv6 addresses only")
	rootCmd.Flags().BoolS("g", "g", false, "Restrict gecos output to real name")
	rootCmd.Flags().BoolS("h", "h", false, "Display remote host")
	rootCmd.Flags().BoolS("k", "k", false, "Disable user accounting database")
	rootCmd.Flags().BoolS("l", "l", false, "Multi-line format")
	rootCmd.Flags().BoolS("m", "m", false, "Prevent matching of user names")
	rootCmd.Flags().BoolS("o", "o", false, "Display office location")
	rootCmd.Flags().BoolS("p", "p", false, "Prevent display of plan files")
	rootCmd.Flags().BoolS("s", "s", false, "Short format")
}
