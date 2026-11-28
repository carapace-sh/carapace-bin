package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createProjectFromExampleCmd = &cobra.Command{
	Use:   "create-project-from-example",
	Short: "Create a project from an example in the ESP Component Registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createProjectFromExampleCmd).Standalone()
	createProjectFromExampleCmd.Flags().String("path", "", "Set the path for the new project")
	carapace.Gen(createProjectFromExampleCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})
	rootCmd.AddCommand(createProjectFromExampleCmd)
}
