package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configPublishCmd = &cobra.Command{
	Use:   "config:publish [--all] [--force] [--] [<name>]",
	Short: "Publish configuration files to your application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configPublishCmd).Standalone()

	configPublishCmd.Flags().Bool("all", false, "Publish all configuration files")
	configPublishCmd.Flags().Bool("force", false, "Overwrite any existing configuration files")
	rootCmd.AddCommand(configPublishCmd)
}
