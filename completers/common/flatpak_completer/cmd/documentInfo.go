package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var documentInfoCmd = &cobra.Command{
	Use:     "document-info [OPTION…] FILE",
	Short:   "Get information about an exported file",
	GroupID: "access",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(documentInfoCmd).Standalone()

	documentInfoCmd.Flags().BoolP("help", "h", false, "Show help options")
	documentInfoCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	documentInfoCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(documentInfoCmd)

	// TODO files from documents command
}
