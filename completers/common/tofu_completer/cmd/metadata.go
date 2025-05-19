package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var metadataCmd = &cobra.Command{
	Use:   "metadata <subcommand> [options] [args]",
	Short: "Metadata related commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(metadataCmd).Standalone()

	rootCmd.AddCommand(metadataCmd)
}
