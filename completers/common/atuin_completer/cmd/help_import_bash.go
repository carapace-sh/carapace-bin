package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_import_bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Import history from the bash history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_import_bashCmd).Standalone()

	help_importCmd.AddCommand(help_import_bashCmd)
}
