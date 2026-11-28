package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createManifestCmd = &cobra.Command{
	Use:   "create-manifest",
	Short: "Create manifest for specified component",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createManifestCmd).Standalone()

	createManifestCmd.Flags().String("component", "", "Name of the component in the project")
	createManifestCmd.Flags().String("path", "", "Path to the component")

	rootCmd.AddCommand(createManifestCmd)

	carapace.Gen(createManifestCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})
}
