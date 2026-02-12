package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var contributorsCmd = &cobra.Command{
	Use:   "contributors",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(contributorsCmd).Standalone()

	contributorsCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(contributorsCmd)
}
