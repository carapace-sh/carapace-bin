package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var balancedCmd = &cobra.Command{
	Use:   "balanced",
	Short: "Apply the balanced profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(balancedCmd).Standalone()

	rootCmd.AddCommand(balancedCmd)
}
