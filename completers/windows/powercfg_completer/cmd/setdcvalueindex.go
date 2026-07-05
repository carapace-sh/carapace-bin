package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setdcvalueindexCmd = &cobra.Command{
	Use:   "setdcvalueindex",
	Short: "set a power setting value for DC power",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setdcvalueindexCmd).Standalone()
	rootCmd.AddCommand(setdcvalueindexCmd)
}
