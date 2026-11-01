package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Builds and pushes multi-arch OpenFaaS container images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(publishCmd).Standalone()
	publishCmd.Flags().StringArrayP("build-arg", "b", nil, "Add a build-arg for Docker (KEY=VALUE)")
	publishCmd.Flags().StringArray("build-label", nil, "Add a label for Docker image (LABEL=VALUE)")
	publishCmd.Flags().StringArrayP("build-option", "o", nil, "Set a build option, e.g. dev")
	publishCmd.Flags().String("builder-public-key", "", "Builder public key as a literal value, or a path to a file containing raw base64 or the JSON response from /publickey")
	publishCmd.Flags().StringArray("copy-extra", nil, "Extra paths that will be copied into the function build context")
	publishCmd.Flags().Bool("debug", false, "Enable debug output when pulling templates")
	publishCmd.Flags().Bool("disable-stack-pull", false, "Disables the template configuration in the stack.yaml")
	publishCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yaml file")
	publishCmd.Flags().StringArray("extra-tag", nil, "Additional extra image tag")
	publishCmd.Flags().String("handler", "", "Directory with handler for function, e.g. handler.js")
	publishCmd.Flags().String("image", "", "Docker image name to build")
	publishCmd.Flags().String("lang", "", "Programming language template")
	publishCmd.Flags().String("name", "", "Name of the deployed function")
	publishCmd.Flags().Bool("no-cache", false, "Do not use Docker's build cache")
	publishCmd.Flags().Bool("overwrite", true, "Overwrite existing templates from the template repository")
	publishCmd.Flags().Int("parallel", 1, "Build in parallel to depth specified.")
	publishCmd.Flags().String("payload-secret", "", "Path to the payload secret file")
	publishCmd.Flags().String("platforms", "linux/amd64", "A set of platforms to publish")
	publishCmd.Flags().Bool("pull", false, "Force a re-pull of base images in template during build, useful for publishing images")
	publishCmd.Flags().Bool("quiet", false, "Perform a quiet build, without showing output from Docker")
	publishCmd.Flags().String("remote-builder", "", "URL to the builder")
	publishCmd.Flags().Bool("reset-qemu", false, "Runs \"docker run multiarch/qemu-user-static --reset -p yes\" to enable multi-arch builds. Compatible with AMD64 machines only.")
	publishCmd.Flags().Bool("shrinkwrap", false, "Just write files to ./build/ folder for shrink-wrapping")
	publishCmd.Flags().Bool("squash", false, "Use Docker's squash flag for smaller images [experimental] ")
	publishCmd.Flags().String("tag", "latest", "Override latest tag on function Docker image, accepts 'digest', 'latest', 'sha', 'branch', or 'describe'")
	rootCmd.AddCommand(publishCmd)

	carapace.Gen(publishCmd).FlagCompletion(carapace.ActionMap{
		"lang":      action.ActionLanguageTemplates(),
		"name":      action.ActionFunctions(),
		"platforms": action.ActionPlatforms(),
		"tag":       carapace.ActionValues("branch", "describe", "digest", "latest", "sha"),
	})
}
