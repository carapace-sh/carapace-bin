package cmd

import (
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Builds, pushes and deploys OpenFaaS function containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upCmd).Standalone()
	upCmd.Flags().StringArray("annotation", nil, "Set one or more annotation (ANNOTATION=VALUE)")
	upCmd.Flags().StringArrayP("build-arg", "b", nil, "Add a build-arg for Docker (KEY=VALUE)")
	upCmd.Flags().StringArray("build-label", nil, "Add a label for Docker image (LABEL=VALUE)")
	upCmd.Flags().StringArrayP("build-option", "o", nil, "Set a build option, e.g. dev")
	upCmd.Flags().String("builder-public-key", "", "Builder public key as a literal value, or a path to a file containing raw base64 or the JSON response from /publickey")
	upCmd.Flags().StringArray("constraint", nil, "Apply a constraint to the function")
	upCmd.Flags().StringArray("copy-extra", nil, "Extra paths that will be copied into the function build context")
	upCmd.Flags().String("cpu-limit", "", "Supply the CPU limit for the function in Mi (when not using a YAML file)")
	upCmd.Flags().String("cpu-request", "", "Supply the CPU request for the function in Mi (when not using a YAML file)")
	upCmd.Flags().Bool("debug", false, "Enable debug output when pulling templates")
	upCmd.Flags().Bool("disable-stack-pull", false, "Disables the template configuration in the stack.yaml")
	upCmd.Flags().StringArrayP("env", "e", nil, "Set one or more environment variables (ENVVAR=VALUE)")
	upCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yaml file")
	upCmd.Flags().String("fprocess", "", "fprocess value to be run as a serverless function by the watchdog")
	upCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	upCmd.Flags().String("handler", "", "Directory with handler for function, e.g. handler.js")
	upCmd.Flags().String("image", "", "Docker image name to build")
	upCmd.Flags().StringArrayP("label", "l", nil, "Set one or more label (LABEL=VALUE)")
	upCmd.Flags().String("lang", "", "Programming language template")
	upCmd.Flags().String("memory-limit", "", "Supply the memory limit for the function in Mi (when not using a YAML file)")
	upCmd.Flags().String("memory-request", "", "Supply the memory request for the function in Mi (when not using a YAML file)")
	upCmd.Flags().String("name", "", "Name of the deployed function")
	upCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	upCmd.Flags().String("network", "", "Name of the network")
	upCmd.Flags().Bool("no-cache", false, "Do not use Docker's build cache")
	upCmd.Flags().Bool("overwrite", true, "Overwrite existing templates from the template repository")
	upCmd.Flags().Int("parallel", 1, "Build in parallel to depth specified.")
	upCmd.Flags().String("payload-secret", "", "Path to the payload secret file")
	upCmd.Flags().String("platforms", "linux/amd64", "Publish for these platforms, when used with --publish")
	upCmd.Flags().Bool("publish", false, "Use faas-cli publish instead of faas-cli build followed by faas-cli push")
	upCmd.Flags().Bool("pull", false, "Force a re-pull of base images in template during build, useful for publishing images")
	upCmd.Flags().Bool("quiet", false, "Perform a quiet build, without showing output from Docker")
	upCmd.Flags().Bool("read-template", true, "Read the function's template")
	upCmd.Flags().Bool("readonly", false, "Force the root container filesystem to be read only")
	upCmd.Flags().String("remote-builder", "", "URL to the builder")
	upCmd.Flags().Bool("replace", false, "Remove and re-create existing function(s)")
	upCmd.Flags().StringArray("secret", nil, "Give the function access to a secure secret")
	upCmd.Flags().Bool("shrinkwrap", false, "Just write files to ./build/ folder for shrink-wrapping")
	upCmd.Flags().Bool("skip-deploy", false, "Skip function deployment")
	upCmd.Flags().Bool("skip-push", false, "Skip pushing function to remote registry")
	upCmd.Flags().Bool("squash", false, "Use Docker's squash flag for smaller images [experimental] ")
	upCmd.Flags().String("tag", "latest", "Override latest tag on function Docker image, accepts 'digest', 'sha', 'branch', or 'describe', or 'latest'")
	upCmd.Flags().Duration("timeout", time.Minute, "Timeout for any HTTP calls made to the OpenFaaS API.")
	upCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	upCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	upCmd.Flags().Bool("update", true, "Perform rolling update on existing function(s)")
	upCmd.Flags().Bool("watch", false, "Watch for changes in files and re-deploy")
	rootCmd.AddCommand(upCmd)

	carapace.Gen(upCmd).FlagCompletion(carapace.ActionMap{
		"copy-extra": carapace.ActionDirectories(),
		"env":        env.ActionNameValues(false),
		"handler":    carapace.ActionDirectories(),
		"lang":       action.ActionLanguageTemplates(),
		"name":       action.ActionFunctions(),
		"namespace":  action.ActionNamespaces(),
		"platforms":  action.ActionPlatforms(),
		"secret":     action.ActionSecrets(),
		"tag":        carapace.ActionValues("branch", "describe", "digest", "latest", "sha"),
	})
}
