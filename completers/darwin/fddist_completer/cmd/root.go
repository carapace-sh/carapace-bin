package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fddist",
	Short: "file descriptor usage distributions",
	Long:  "https://man.freebsd.org/cgi/man.cgi?fddist",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("r", "r", false, "Reads only")
	rootCmd.Flags().BoolS("w", "w", false, "Writes only")
}
