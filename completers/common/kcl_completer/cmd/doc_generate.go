package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var doc_generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates documents from code and examples",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doc_generateCmd).Standalone()

	doc_generateCmd.Flags().Bool("escape-html", false, "Whether to escape html symbols when the output format is markdown")
	doc_generateCmd.Flags().String("file-path", "", "Relative or absolute path to the KCL package root")
	doc_generateCmd.Flags().String("format", "markdown", "The document format to generate")
	doc_generateCmd.Flags().Bool("ignore-deprecated", false, "Do not generate documentation for deprecated schemas")
	doc_generateCmd.Flags().String("target", "", "The target directory where the docs/ folder will be created")
	doc_generateCmd.Flags().String("template", "", "The template directory based on the KCL package root")
	docCmd.AddCommand(doc_generateCmd)

	carapace.Gen(doc_generateCmd).FlagCompletion(carapace.ActionMap{
		"file-path": carapace.ActionDirectories(),
		"format":    carapace.ActionValues("markdown", "html", "openapi"),
		"target":    carapace.ActionDirectories(),
		"template":  carapace.ActionDirectories(),
	})
}
