package cmd

import (
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
	upCmd.Flags().StringArray("constraint", nil, "Apply a constraint to the function")
	upCmd.Flags().StringArray("copy-extra", nil, "Extra paths that will be copied into the function build context")
	upCmd.Flags().Bool("disable-stack-pull", false, "Disables the template configuration in the stack.yml")
	upCmd.Flags().StringArrayP("env", "e", nil, "Set one or more environment variables (ENVVAR=VALUE)")
	upCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yml file")
	upCmd.Flags().String("fprocess", "", "fprocess value to be run as a serverless function by the watchdog")
	upCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	upCmd.Flags().String("handler", "", "Directory with handler for function, e.g. handler.js")
	upCmd.Flags().String("image", "", "Docker image name to build")
	upCmd.Flags().StringArrayP("label", "l", nil, "Set one or more label (LABEL=VALUE)")
	upCmd.Flags().String("lang", "", "Programming language template")
	upCmd.Flags().String("name", "", "Name of the deployed function")
	upCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	upCmd.Flags().String("network", "", "Name of the network")
	upCmd.Flags().Bool("no-cache", false, "Do not use Docker's build cache")
	upCmd.Flags().Int("parallel", 1, "Build in parallel to depth specified.")
	upCmd.Flags().Bool("quiet", false, "Perform a quiet build, without showing output from Docker")
	upCmd.Flags().Bool("read-template", true, "Read the function's template")
	upCmd.Flags().Bool("readonly", false, "Force the root container filesystem to be read only")
	upCmd.Flags().Bool("replace", false, "Remove and re-create existing function(s)")
	upCmd.Flags().StringArray("secret", nil, "Give the function access to a secure secret")
	upCmd.Flags().Bool("shrinkwrap", false, "Just write files to ./build/ folder for shrink-wrapping")
	upCmd.Flags().Bool("skip-deploy", false, "Skip function deployment")
	upCmd.Flags().Bool("skip-push", false, "Skip pushing function to remote registry")
	upCmd.Flags().Bool("squash", false, "Use Docker's squash flag for smaller images [experimental] ")
	upCmd.Flags().String("tag", "latest", "Override latest tag on function Docker image, accepts 'latest', 'sha', 'branch', or 'describe'")
	upCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	upCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	upCmd.Flags().Bool("update", true, "Perform rolling update on existing function(s)")
	rootCmd.AddCommand(upCmd)

	carapace.Gen(upCmd).FlagCompletion(carapace.ActionMap{
		"copy-extra": carapace.ActionDirectories(),
		"env":        env.ActionNameValues(false),
		"handler":    carapace.ActionDirectories(),
		"lang":       action.ActionLanguageTemplates(),
		"name":       action.ActionFunctions(),
		"namespace":  action.ActionNamespaces(),
		"tag":        carapace.ActionValues("latest", "sha", "branch", "describe"),
	})
}
