package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize shell integration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_initCmd).Standalone()

	ai_initCmd.Flags().BoolP("help", "h", false, "Print help")
	aiCmd.AddCommand(ai_initCmd)

	carapace.Gen(ai_initCmd).PositionalCompletion(
		carapace.ActionValues("zsh", "bash", "fish", "auto"),
	)
}
