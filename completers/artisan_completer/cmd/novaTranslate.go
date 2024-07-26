package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaTranslateCmd = &cobra.Command{
	Use:   "nova:translate [--force] [--] <language>",
	Short: "Create translation files for Nova",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaTranslateCmd).Standalone()

	novaTranslateCmd.Flags().Bool("force", false, "Overwrite any existing files")
	rootCmd.AddCommand(novaTranslateCmd)
}
