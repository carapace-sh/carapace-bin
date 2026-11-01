package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds OpenFaaS function containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().StringArrayP("build-arg", "b", nil, "Add a build-arg for Docker (KEY=VALUE)")
	buildCmd.Flags().StringArray("build-label", nil, "Add a label for Docker image (LABEL=VALUE)")
	buildCmd.Flags().StringArrayP("build-option", "o", nil, "Set a build option, e.g. dev")
	buildCmd.Flags().StringArray("copy-extra", nil, "Extra paths that will be copied into the function build context")
	buildCmd.Flags().Bool("debug", false, "Enable debug output when pulling templates")
	buildCmd.Flags().Bool("disable-stack-pull", false, "Disables the template configuration in the stack.yaml")
	buildCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yaml file")
	buildCmd.Flags().String("handler", "", "Directory with handler for function, e.g. handler.js")
	buildCmd.Flags().String("image", "", "Docker image name to build")
	buildCmd.Flags().String("lang", "", "Programming language template")
	buildCmd.Flags().String("name", "", "Name of the deployed function")
	buildCmd.Flags().Bool("no-cache", false, "Do not use Docker's build cache")
	buildCmd.Flags().Bool("overwrite", true, "Overwrite existing templates from the template repository")
	buildCmd.Flags().Int("parallel", 1, "Build in parallel to depth specified.")
	buildCmd.Flags().Bool("pull", false, "Force a re-pull of base images in template during build, useful for publishing images")
	buildCmd.Flags().Bool("quiet", false, "Perform a quiet build, without showing output from Docker")
	buildCmd.Flags().Bool("shrinkwrap", false, "Just write files to ./build/ folder for shrink-wrapping")
	buildCmd.Flags().Bool("squash", false, "Use Docker's squash flag for smaller images [experimental] ")
	buildCmd.Flags().String("tag", "latest", "Override latest tag on function Docker image, accepts 'digest', 'sha', 'branch', or 'describe', or 'latest'")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"copy-extra": carapace.ActionDirectories(),
		"handler":    carapace.ActionDirectories(),
		"lang":       action.ActionLanguageTemplates(),
		"tag":        carapace.ActionValues("branch", "describe", "digest", "latest", "sha"),
	})
}
