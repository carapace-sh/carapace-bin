package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().String("id", "", "Deployment ID to pull environment variables from")
	buildCmd.Flags().String("output", "", "Directory where built assets will be written to")
	buildCmd.Flags().Bool("prod", false, "Build a production deployment")
	buildCmd.Flags().String("project", "", "Project name or ID")
	buildCmd.Flags().Bool("standalone", false, "Create a standalone build with all dependencies inlined")
	buildCmd.Flags().String("target", "", "Specify the target environment")
	buildCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt")

	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"project": action.ActionProjects(buildCmd),
		"target":  carapace.ActionValues("preview", "prod"),
	})
}
