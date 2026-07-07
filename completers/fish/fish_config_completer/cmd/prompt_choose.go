package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var promptChooseCmd = &cobra.Command{
	Use:   "choose",
	Short: "Load a sample prompt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(promptChooseCmd).Standalone()

	promptCmd.AddCommand(promptChooseCmd)
}
