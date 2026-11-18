package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_initCmd = &cobra.Command{
	Use:    "init",
	Short:  "This command is no longer necessary. If you have it in your shell init file, feel free to remove it",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_initCmd).Standalone()

	ai_initCmd.Flags().BoolP("help", "h", false, "Print help")
	aiCmd.AddCommand(ai_initCmd)
}
