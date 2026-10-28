package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var localRunCmd = &cobra.Command{
	Use:   "local-run",
	Short: "Start a function with docker for local testing (experimental feature)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(localRunCmd).Standalone()

	localRunCmd.Flags().Bool("build", true, "Build function prior to local-run")
	localRunCmd.Flags().StringArrayP("build-arg", "b", nil, "Add a build-arg for Docker (KEY=VALUE)")
	localRunCmd.Flags().StringArray("build-label", nil, "Add a label for Docker image (LABEL=VALUE)")
	localRunCmd.Flags().StringArrayP("build-option", "o", nil, "Set a build option, e.g. dev")
	localRunCmd.Flags().StringArray("copy-extra", nil, "Extra paths that will be copied into the function build context")
	localRunCmd.Flags().Bool("debug", false, "Enable debug output when pulling templates")
	localRunCmd.Flags().Bool("disable-stack-pull", false, "Disables the template configuration in the stack.yaml")
	localRunCmd.Flags().StringToStringP("env", "e", nil, "additional environment variables (ENVVAR=VALUE), use this to experiment with different values for your function")
	localRunCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yaml file")
	localRunCmd.Flags().String("handler", "", "Directory with handler for function, e.g. handler.js")
	localRunCmd.Flags().String("image", "", "Docker image name to build")
	localRunCmd.Flags().String("lang", "", "Programming language template")
	localRunCmd.Flags().String("name", "", "Name of the deployed function")
	localRunCmd.Flags().String("network", "", "connect function to an existing network, use 'host' to access other process already running on localhost. When using this, '--port' is ignored, if you have port collisions, you may change the port using '-e port=NEW_PORT'")
	localRunCmd.Flags().Bool("no-cache", false, "Do not use Docker's build cache")
	localRunCmd.Flags().Bool("overwrite", true, "Overwrite existing templates from the template repository")
	localRunCmd.Flags().Int("parallel", 1, "Build in parallel to depth specified.")
	localRunCmd.Flags().IntP("port", "p", 8080, "port to bind the function to, set to \"0\" to use a random port")
	localRunCmd.Flags().Bool("print", false, "Print the docker command instead of running it")
	localRunCmd.Flags().Bool("pull", false, "Force a re-pull of base images in template during build, useful for publishing images")
	localRunCmd.Flags().Bool("quiet", false, "Perform a quiet build, without showing output from Docker")
	localRunCmd.Flags().Bool("shrinkwrap", false, "Just write files to ./build/ folder for shrink-wrapping")
	localRunCmd.Flags().Bool("squash", false, "Use Docker's squash flag for smaller images [experimental] ")
	localRunCmd.Flags().String("tag", "latest", "Override latest tag on function Docker image, accepts 'digest', 'sha', 'branch', or 'describe', or 'latest'")
	localRunCmd.Flags().Bool("watch", false, "Watch for changes in files and re-deploy")
	rootCmd.AddCommand(localRunCmd)

	carapace.Gen(localRunCmd).FlagCompletion(carapace.ActionMap{
		"copy-extra": carapace.ActionDirectories(),
		"env":        env.ActionNameValues(false),
		"handler":    carapace.ActionDirectories(),
		"lang":       action.ActionLanguageTemplates(),
		"tag":        carapace.ActionValues("branch", "describe", "digest", "latest", "sha"),
	})
}
