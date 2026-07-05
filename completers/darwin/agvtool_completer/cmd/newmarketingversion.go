package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var newMarketingVersionCmd = &cobra.Command{
	Use:   "new-marketing-version",
	Short: "Set the marketing version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newMarketingVersionCmd).Standalone()
	rootCmd.AddCommand(newMarketingVersionCmd)
}
