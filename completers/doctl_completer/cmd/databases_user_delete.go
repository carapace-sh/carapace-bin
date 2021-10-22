package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_user_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a database user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_user_deleteCmd).Standalone()
	databases_user_deleteCmd.Flags().BoolP("force", "f", false, "Delete the user without a confirmation prompt")
	databases_userCmd.AddCommand(databases_user_deleteCmd)
}
