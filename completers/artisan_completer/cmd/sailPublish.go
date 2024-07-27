package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sailPublishCmd = &cobra.Command{
	Use:   "sail:publish",
	Short: "Publish the Laravel Sail Docker files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sailPublishCmd).Standalone()

	rootCmd.AddCommand(sailPublishCmd)
}
