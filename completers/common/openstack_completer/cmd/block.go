package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockCmd).Standalone()

	rootCmd.AddCommand(blockCmd)
}
