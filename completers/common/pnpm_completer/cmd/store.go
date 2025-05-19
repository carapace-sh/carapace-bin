package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:     "store",
	Short:   "Reads and performs actions on pnpm store that is on the current filesystem",
	GroupID: "store",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storeCmd).Standalone()

	rootCmd.AddCommand(storeCmd)
}
