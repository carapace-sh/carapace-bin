package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_help_bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Import history from the bash history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_help_bashCmd).Standalone()

	import_helpCmd.AddCommand(import_help_bashCmd)
}
