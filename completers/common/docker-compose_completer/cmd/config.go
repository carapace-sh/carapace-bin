package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config [OPTIONS] [SERVICE...]",
	Short: "Parse, resolve and render compose file in canonical format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().Bool("environment", false, "Print environment used for interpolation.")
	configCmd.Flags().String("format", "", "Format the output. Values: [yaml | json]")
	configCmd.Flags().String("hash", "", "Print the service config hash, one per line.")
	configCmd.Flags().Bool("images", false, "Print the image names, one per line.")
	configCmd.Flags().Bool("lock-image-digests", false, "Produces an override file with image digests")
	configCmd.Flags().Bool("models", false, "Print the model names, one per line.")
	configCmd.Flags().Bool("networks", false, "Print the network names, one per line.")
	configCmd.Flags().Bool("no-consistency", false, "Don't check model consistency - warning: may produce invalid Compose output")
	configCmd.Flags().Bool("no-env-resolution", false, "Don't resolve service env files")
	configCmd.Flags().Bool("no-interpolate", false, "Don't interpolate environment variables")
	configCmd.Flags().Bool("no-normalize", false, "Don't normalize compose model")
	configCmd.Flags().Bool("no-path-resolution", false, "Don't resolve file paths")
	configCmd.Flags().StringP("output", "o", "", "Save to file (default to stdout)")
	configCmd.Flags().Bool("profiles", false, "Print the profile names, one per line.")
	configCmd.Flags().BoolP("quiet", "q", false, "Only validate the configuration, don't print anything")
	configCmd.Flags().Bool("resolve-image-digests", false, "Pin image tags to digests")
	configCmd.Flags().Bool("services", false, "Print the service names, one per line.")
	configCmd.Flags().Bool("variables", false, "Print model variables and default values.")
	configCmd.Flags().Bool("volumes", false, "Print the volume names, one per line.")
	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("yaml", "json"),
		"hash":   action.ActionServices(configCmd),
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(configCmd).PositionalCompletion(
		action.ActionServices(configCmd).FilterArgs(),
	)
}
