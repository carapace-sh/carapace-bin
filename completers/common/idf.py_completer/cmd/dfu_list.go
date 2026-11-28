package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dfuListCmd = &cobra.Command{
	Use:   "dfu-list",
	Short: "List DFU capable devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dfuListCmd).Standalone()
	rootCmd.AddCommand(dfuListCmd)
}
