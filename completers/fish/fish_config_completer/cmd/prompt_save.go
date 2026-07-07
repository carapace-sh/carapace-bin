package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var promptSaveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save the current prompt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(promptSaveCmd).Standalone()

	promptCmd.AddCommand(promptSaveCmd)
}
