package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deployKey_deleteCmd = &cobra.Command{
	Use:   "delete <key-id>",
	Short: "Deletes a single deploy key specified by the ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployKey_deleteCmd).Standalone()

	deployKeyCmd.AddCommand(deployKey_deleteCmd)

	// TODO positional completion
}
