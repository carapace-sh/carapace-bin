package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the golangci-lint version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().Bool("debug", false, "Add build information")
	versionCmd.Flags().Bool("json", false, "Display as JSON")
	versionCmd.Flags().Bool("short", false, "Display only the version number")
	rootCmd.AddCommand(versionCmd)
}
