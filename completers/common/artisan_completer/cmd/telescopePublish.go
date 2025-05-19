package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var telescopePublishCmd = &cobra.Command{
	Use:   "telescope:publish [--force]",
	Short: "Publish all of the Telescope resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(telescopePublishCmd).Standalone()

	telescopePublishCmd.Flags().Bool("force", false, "Overwrite any existing files")
	rootCmd.AddCommand(telescopePublishCmd)
}
