package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:     "cache",
	Short:   "manage cache",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cacheCmd).Standalone()

	rootCmd.AddCommand(cacheCmd)
}
