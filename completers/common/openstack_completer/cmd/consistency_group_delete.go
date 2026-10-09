package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consistency_group_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete consistency group(s).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consistency_group_deleteCmd).Standalone()

	consistency_group_deleteCmd.Flags().Bool("force", false, "Allow delete in state other than error or available")
	consistency_groupCmd.AddCommand(consistency_group_deleteCmd)
}
