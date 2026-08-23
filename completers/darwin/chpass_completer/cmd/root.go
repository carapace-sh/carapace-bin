package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chpass",
	Short: "add or change user database information",
	Long:  "https://keith.github.io/xcode-manpages/chpass.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("l", "l", "", "Location of the directory node")
	rootCmd.Flags().StringS("s", "s", "", "Shell to change to")
	rootCmd.Flags().StringS("u", "u", "", "User name to use when authenticating")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"s": carapace.ActionFiles(),
	})
}
