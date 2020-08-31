package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Validate and view the Compose file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	configCmd.Flags().String("hash", "", "Print the service config hash, one per line.")
	configCmd.Flags().Bool("no-interpolate", false, "Don't interpolate environment variables")
	configCmd.Flags().BoolP("quiet", "q", false, "Only validate the configuration, don't print")
	configCmd.Flags().Bool("resolve-image-digests", false, "Pin image tags to digests.")
	configCmd.Flags().Bool("services", false, "Print the service names, one per line.")
	configCmd.Flags().Bool("volumes", false, "Print the volume names, one per line.")
	rootCmd.AddCommand(configCmd)
}
