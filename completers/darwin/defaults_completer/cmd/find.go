package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Search for a key in all domains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findCmd).Standalone()
	rootCmd.AddCommand(findCmd)
}
