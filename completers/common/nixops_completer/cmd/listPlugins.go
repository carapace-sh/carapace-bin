package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listPluginsCmd = &cobra.Command{
	Use:   "list-plugins",
	Short: "list available nixops plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listPluginsCmd).Standalone()

	listPluginsCmd.Flags().Bool("debug", false, "Enable debug output")
	listPluginsCmd.Flags().BoolP("verbose", "v", false, "Provide extra plugin information")

	rootCmd.AddCommand(listPluginsCmd)
}
