package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(usageCmd).Standalone()

	rootCmd.AddCommand(usageCmd)
}
