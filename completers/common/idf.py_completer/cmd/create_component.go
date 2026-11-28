package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createComponentCmd = &cobra.Command{
	Use:   "create-component",
	Short: "Create a new component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createComponentCmd).Standalone()

	rootCmd.AddCommand(createComponentCmd)
}
