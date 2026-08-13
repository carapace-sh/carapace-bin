package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vercel",
	Short: "Develop. Preview. Ship.",
	Long:  "https://vercel.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().String("api", "", "(deprecated)")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Debug mode (default off)")
	rootCmd.PersistentFlags().String("cwd", "", "Sets the current working directory")
	rootCmd.PersistentFlags().StringP("global-config", "Q", "", "Path to the global `.vercel` directory")
	rootCmd.PersistentFlags().StringP("local-config", "A", "", "Path to the local `vercel.json` file")
	rootCmd.PersistentFlags().Bool("no-color", false, "No color mode (default off)")
	rootCmd.PersistentFlags().Bool("non-interactive", false, "Run without interactive prompts")
	rootCmd.PersistentFlags().StringP("scope", "S", "", "Set a custom scope")
	rootCmd.PersistentFlags().StringP("token", "t", "", "Login token")
	rootCmd.PersistentFlags().StringP("team", "T", "", "(deprecated)")

	rootCmd.Flags().Bool("archive", false, "Compress the deployment code into an archive before uploading it")
	rootCmd.Flags().StringSliceP("build-env", "b", nil, "Similar to `--env` but for build time only.")
	rootCmd.Flags().BoolP("confirm", "c", false, "(deprecated)")
	rootCmd.Flags().Bool("dry", false, "Inspect detected framework preset without deploying")
	rootCmd.Flags().StringSliceP("env", "e", nil, "Include an env var during run time (e.g.: `-e KEY=value`). Can appear many times.")
	rootCmd.Flags().BoolP("force", "f", false, "Force a new deployment even if nothing has changed")
	rootCmd.Flags().String("guidance", "", "Receive command suggestions once deployment is complete")
	rootCmd.Flags().BoolP("logs", "l", false, "Print the build logs")
	rootCmd.Flags().StringSliceP("meta", "m", nil, "Add metadata for the deployment (e.g.: `-m KEY=value`). Can appear many times.")
	rootCmd.Flags().BoolP("no-clipboard", "C", false, "Do not attempt to copy URL to clipboard")
	rootCmd.Flags().Bool("no-wait", false, "Don't wait for the deployment to finish")
	rootCmd.Flags().BoolP("platform-version", "V", false, "Set the platform version to deploy to")
	rootCmd.Flags().Bool("prebuilt", false, "Use in combination with `vc build`. Deploy an existing build")
	rootCmd.Flags().Bool("prod", false, "Create a production deployment")
	rootCmd.Flags().String("project", "", "Project name or ID")
	rootCmd.Flags().BoolP("public", "p", false, "Deployment is public (`/_src` is exposed)")
	rootCmd.Flags().String("regions", "", "Set default regions to enable the deployment on")
	rootCmd.Flags().Bool("skip-domain", false, "Disable automatic promotion of domains")
	rootCmd.Flags().String("target", "", "Specify the target deployment environment")
	rootCmd.Flags().Bool("with-cache", false, "Retain build cache when using \"--force\"")
	rootCmd.Flags().BoolP("yes", "y", false, "Use default options to skip all prompts")

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Output usage information")
	rootCmd.Flags().BoolP("version", "v", false, "Output the version number")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"build-env":     env.ActionNameValues(false),
		"env":           env.ActionNameValues(false),
		"global-config": carapace.ActionDirectories(),
		"local-config":  carapace.ActionFiles(),
		"regions":       action.ActionRegions().UniqueList(","),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
