package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports a file into an environment in an existing workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_importCmd).Standalone()

	helpCmd.AddCommand(help_importCmd)
}
