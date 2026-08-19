package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(EditCmd).Standalone()
	rootCmd.AddCommand(EditCmd)
}
