package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var preview_prompterCmd = &cobra.Command{
	Use:   "prompter [prompt type]",
	Short: "Execute a test program to preview the prompter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(preview_prompterCmd).Standalone()

	previewCmd.AddCommand(preview_prompterCmd)

	carapace.Gen(preview_prompterCmd).PositionalCompletion(
		carapace.ActionValues(
			"select",
			"multi-select",
			"input",
			"password",
			"confirm",
			"auth-token",
			"confirm-deletion",
			"input-hostname",
			"markdown-editor",
		),
	)
}
