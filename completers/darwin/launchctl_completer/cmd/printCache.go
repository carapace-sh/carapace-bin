package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var printCacheCmd = &cobra.Command{
	Use:   "print-cache",
	Short: "Print information about the service cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(printCacheCmd).Standalone()
	rootCmd.AddCommand(printCacheCmd)
}
