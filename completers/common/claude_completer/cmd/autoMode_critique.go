package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoMode_critiqueCmd = &cobra.Command{
	Use:   "critique",
	Short: "Get AI feedback on your custom auto mode rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoMode_critiqueCmd).Standalone()

	autoMode_critiqueCmd.Flags().BoolP("help", "h", false, "Display help for command")
	autoMode_critiqueCmd.Flags().String("model", "", "Override which model is used")
	autoModeCmd.AddCommand(autoMode_critiqueCmd)

	carapace.Gen(autoMode_critiqueCmd).FlagCompletion(carapace.ActionMap{
		"model": carapace.ActionValues("fable", "haiku", "opus", "sonnet"),
	})
}
