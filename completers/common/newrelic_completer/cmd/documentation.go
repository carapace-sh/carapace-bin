package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var documentationCmd = &cobra.Command{
	Use:   "documentation",
	Short: "Generate CLI documentation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(documentationCmd).Standalone()
	documentationCmd.Flags().String("format", "markdown", "Documentation format [markdown, manpage] default 'markdown'")
	documentationCmd.Flags().StringP("outputDir", "o", "", "Output directory for generated documentation")
	rootCmd.AddCommand(documentationCmd)

	carapace.Gen(documentationCmd).FlagCompletion(carapace.ActionMap{
		"format":    carapace.ActionValues("markdown", "manpage"),
		"outputDir": carapace.ActionDirectories(),
	})
}
