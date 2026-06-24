package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var printCacheCmd = &cobra.Command{
	Use:   "print-cache",
	Short: "Prints information about the service cache",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(printCacheCmd).Standalone()
	rootCmd.AddCommand(printCacheCmd)
}
