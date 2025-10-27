package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generateDocsCmd = &cobra.Command{
	Use:    "generate-docs",
	Short:  "Populates the specified folder with documentation in markdown about minikube",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateDocsCmd).Standalone()

	generateDocsCmd.Flags().String("code-path", "", "The path on the file system where the error code docs in markdown need to be saved")
	generateDocsCmd.Flags().String("path", "", "The path on the file system where the docs in markdown need to be saved")
	generateDocsCmd.Flags().String("test-path", "", "The path on the file system where the testing docs in markdown need to be saved")
	rootCmd.AddCommand(generateDocsCmd)

	carapace.Gen(generateDocsCmd).FlagCompletion(carapace.ActionMap{
		"code-path": carapace.ActionDirectories(),
		"path":      carapace.ActionDirectories(),
		"test-path": carapace.ActionDirectories(),
	})
}
