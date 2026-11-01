package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var AddDriverCmd = &cobra.Command{
	Use:   "Add-Driver",
	Short: "add third-party driver packages to an offline image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(AddDriverCmd).Standalone()
	rootCmd.AddCommand(AddDriverCmd)
}
