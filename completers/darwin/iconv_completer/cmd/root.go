package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iconv",
	Short: "codeset conversion utility",
	Long:  "https://man.freebsd.org/cgi/man.cgi?iconv",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "Prevent output of invalid characters")
	rootCmd.Flags().StringS("f", "f", "", "Source codeset name")
	rootCmd.Flags().BoolS("l", "l", false, "List available codeset names")
	rootCmd.Flags().BoolS("s", "s", false, "Silent")
	rootCmd.Flags().StringS("t", "t", "", "Target codeset name")
}
