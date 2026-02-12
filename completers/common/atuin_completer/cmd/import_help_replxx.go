package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_help_replxxCmd = &cobra.Command{
	Use:   "replxx",
	Short: "Import history from the replxx history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_help_replxxCmd).Standalone()

	import_helpCmd.AddCommand(import_help_replxxCmd)
}
