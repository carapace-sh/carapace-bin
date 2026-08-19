package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(DeleteCmd).Standalone()
	rootCmd.AddCommand(DeleteCmd)
}
