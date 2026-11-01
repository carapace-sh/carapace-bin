package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "klist",
	Short: "list cached Kerberos tickets",
	Long:  "https://man.freebsd.org/cgi/man.cgi?klist",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("5", "5", false, "Show v5 tickets")
	rootCmd.Flags().BoolS("T", "T", false, "Tokens")
	rootCmd.Flags().StringS("c", "c", "", "Cache name")
	rootCmd.Flags().BoolS("f", "f", false, "Include ticket flags")
	rootCmd.Flags().BoolS("l", "l", false, "List caches")
	rootCmd.Flags().BoolS("s", "s", false, "Test")
	rootCmd.Flags().BoolS("t", "t", false, "Test")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionFiles(),
	})
}
