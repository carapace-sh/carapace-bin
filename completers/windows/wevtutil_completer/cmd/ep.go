package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var epCmd = &cobra.Command{
	Use:   "ep",
	Short: "enumerate event publishers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(epCmd).Standalone()
	rootCmd.AddCommand(epCmd)
}
