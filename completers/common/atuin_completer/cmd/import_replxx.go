package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_replxxCmd = &cobra.Command{
	Use:   "replxx",
	Short: "Import history from the replxx history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_replxxCmd).Standalone()

	import_replxxCmd.Flags().BoolP("help", "h", false, "Print help")
	importCmd.AddCommand(import_replxxCmd)
}
