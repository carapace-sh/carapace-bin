package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var DestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(DestroyCmd).Standalone()
	rootCmd.AddCommand(DestroyCmd)
}
