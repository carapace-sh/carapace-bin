package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Import history from the bash history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_bashCmd).Standalone()

	import_bashCmd.Flags().BoolP("help", "h", false, "Print help")
	importCmd.AddCommand(import_bashCmd)
}
