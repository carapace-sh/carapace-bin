package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bisect_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show long usage description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_helpCmd).Standalone()

	bisectCmd.AddCommand(bisect_helpCmd)
}
