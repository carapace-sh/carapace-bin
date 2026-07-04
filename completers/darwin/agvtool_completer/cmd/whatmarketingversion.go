package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var whatMarketingVersionCmd = &cobra.Command{
	Use:   "what-marketing-version",
	Short: "Print the marketing version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whatMarketingVersionCmd).Standalone()
	whatMarketingVersionCmd.Flags().Bool("terse", false, "Print only the version number")
	whatMarketingVersionCmd.Flags().Bool("terse1", false, "Print only the first version found")
	rootCmd.AddCommand(whatMarketingVersionCmd)
}
