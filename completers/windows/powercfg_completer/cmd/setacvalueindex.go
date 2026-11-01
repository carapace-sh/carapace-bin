package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setacvalueindexCmd = &cobra.Command{
	Use:   "setacvalueindex",
	Short: "set a power setting value for AC power",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setacvalueindexCmd).Standalone()
	rootCmd.AddCommand(setacvalueindexCmd)
}
