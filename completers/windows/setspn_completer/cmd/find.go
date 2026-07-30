package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "find duplicate SPNs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findCmd).Standalone()
	rootCmd.AddCommand(findCmd)
}
