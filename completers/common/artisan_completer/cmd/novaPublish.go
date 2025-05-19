package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaPublishCmd = &cobra.Command{
	Use:   "nova:publish [--force]",
	Short: "Publish all of the Nova resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaPublishCmd).Standalone()

	novaPublishCmd.Flags().Bool("force", false, "Overwrite any existing files")
	rootCmd.AddCommand(novaPublishCmd)
}
