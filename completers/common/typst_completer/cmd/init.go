package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new project from a template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().String("package-path", "", "Custom path to local packages, defaults to system-dependent location")
	initCmd.Flags().String("package-cache-path", "", "Custom path to package cache, defaults to system-dependent location")
	initCmd.Flags().BoolP("help", "h", false, "Print help (see a summary with '-h')")

	rootCmd.AddCommand(initCmd)

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"package-path":       carapace.ActionDirectories(),
		"package-cache-path": carapace.ActionDirectories(),
	})

	// TODO: PositionalCompletion for templates
}
