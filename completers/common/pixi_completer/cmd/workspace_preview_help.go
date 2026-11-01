package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_preview_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print help for subcommands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_preview_helpCmd).Standalone()
	workspace_previewCmd.AddCommand(workspace_preview_helpCmd)
}
