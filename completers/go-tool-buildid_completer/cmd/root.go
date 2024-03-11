package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-tool-buildid",
	Short: "Buildid displays or updates the build ID stored in a Go package or binary.",
	Long:  "https://pkg.go.dev/cmd/buildid",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("w", "w", false, "write build ID")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
