package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tailspin",
	Short: "configure, save and print tailspin output",
	Long:  "https://man.freebsd.org/cgi/man.cgi?tailspin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.Batch(
		carapace.ActionValues("info", "enable", "disable", "set", "reset", "save", "augment", "stat"),
	).ToA())
}
