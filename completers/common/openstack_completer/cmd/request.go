package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(requestCmd).Standalone()

	rootCmd.AddCommand(requestCmd)
}
