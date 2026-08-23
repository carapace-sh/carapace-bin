package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stdbuf",
	Short: "change standard streams initial buffering",
	Long:  "https://man.freebsd.org/cgi/man.cgi?stdbuf",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("e", "e", "", "Set initial buffering of standard error")
	rootCmd.Flags().StringS("i", "i", "", "Set initial buffering of standard input")
	rootCmd.Flags().StringS("o", "o", "", "Set initial buffering of standard output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"e": carapace.ActionValues("0", "L", "B"),
		"i": carapace.ActionValues("0", "L", "B"),
		"o": carapace.ActionValues("0", "L", "B"),
	})
}
