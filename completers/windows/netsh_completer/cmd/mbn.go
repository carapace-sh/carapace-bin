package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mbnCmd = &cobra.Command{
	Use:   "mbn",
	Short: "mobile broadband network configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mbnCmd).Standalone()
	rootCmd.AddCommand(mbnCmd)
}
