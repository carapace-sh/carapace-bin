package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Manage the mise cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cacheCmd).Standalone()

	cacheCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(cacheCmd)
}
