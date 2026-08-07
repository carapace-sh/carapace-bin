package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securefile_removeCmd = &cobra.Command{
	Use:     "remove [<id> | --id <id> | --name <name>] [flags]",
	Short:   "Remove a secure file from a project.",
	Aliases: []string{"rm", "delete"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securefile_removeCmd).Standalone()

	securefile_removeCmd.Flags().String("id", "", "ID of the secure file to remove.")
	securefile_removeCmd.Flags().String("name", "", "Name of the secure file to remove.")
	securefile_removeCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt.")
	securefileCmd.AddCommand(securefile_removeCmd)

	// TODO complete file ids
}
