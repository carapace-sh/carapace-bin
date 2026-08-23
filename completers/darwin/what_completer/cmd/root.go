package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "what",
	Short: "show what versions of object modules were used to construct a file",
	Long:  "https://man.freebsd.org/cgi/man.cgi?what",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("q", "q", false, "Only output the match text")
	rootCmd.Flags().BoolS("s", "s", false, "Stop searching after the first match")
}
