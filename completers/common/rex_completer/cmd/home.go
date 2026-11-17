package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var homeCmd = &cobra.Command{
	Use:   "home",
	Short: "Show or choose your home server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(homeCmd).Standalone()

	homeCmd.Flags().Bool("json", false, "output the home and your identity as JSON")

	rootCmd.AddCommand(homeCmd)
}
