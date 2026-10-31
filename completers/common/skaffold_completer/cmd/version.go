package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().StringP("output", "o", "{{.Version}}", "Format output with go-template. For full struct documentation, see https://godoc.org/github.com/GoogleContainerTools/skaffold/v2/pkg/skaffold/version#Info")
	applyCommonFlagCompletion(versionCmd)
	rootCmd.AddCommand(versionCmd)
}
