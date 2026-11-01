package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Manipulate packages cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cacheCmd).Standalone()
	cacheCmd.PersistentFlags().String("cache", "", "location of npm's cache directory")

	rootCmd.AddCommand(cacheCmd)
}
