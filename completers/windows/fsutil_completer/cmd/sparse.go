package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sparseCmd = &cobra.Command{
	Use:   "sparse",
	Short: "sparse file control",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparseCmd).Standalone()
	rootCmd.AddCommand(sparseCmd)
}
