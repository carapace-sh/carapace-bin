package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_import_replxxCmd = &cobra.Command{
	Use:   "replxx",
	Short: "Import history from the replxx history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_import_replxxCmd).Standalone()

	help_importCmd.AddCommand(help_import_replxxCmd)
}
