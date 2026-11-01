package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nlmCmd = &cobra.Command{
	Use:   "nlm",
	Short: "Network Location Manager configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nlmCmd).Standalone()
	rootCmd.AddCommand(nlmCmd)
}
