package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bannerCmd = &cobra.Command{
	Use:   "banner",
	Short: "print installation banner",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bannerCmd).Standalone()

	bannerCmd.Flags().BoolS("v", "v", false, "verbosity")

	rootCmd.AddCommand(bannerCmd)
}
