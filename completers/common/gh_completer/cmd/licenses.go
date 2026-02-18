package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licensesCmd = &cobra.Command{
	Use:   "licenses",
	Short: "View third-party license information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licensesCmd).Standalone()

	rootCmd.AddCommand(licensesCmd)
}
