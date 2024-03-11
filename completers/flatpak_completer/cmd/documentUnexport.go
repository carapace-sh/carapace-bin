package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var documentUnexportCmd = &cobra.Command{
	Use:     "document-unexport [OPTION…] FILE",
	Short:   "Unexport a file to apps",
	GroupID: "access",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(documentUnexportCmd).Standalone()

	documentUnexportCmd.Flags().Bool("doc-id", false, "Specify the document ID")
	documentUnexportCmd.Flags().BoolP("help", "h", false, "Show help options")
	documentUnexportCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	documentUnexportCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(documentUnexportCmd)

	// TODO files from documents command
}
