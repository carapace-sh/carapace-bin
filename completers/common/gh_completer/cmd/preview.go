package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var previewCmd = &cobra.Command{
	Use:   "preview <command>",
	Short: "Execute previews for gh features",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(previewCmd).Standalone()

	rootCmd.AddCommand(previewCmd)
}
