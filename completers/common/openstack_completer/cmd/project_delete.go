package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var project_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete project(s).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_deleteCmd).Standalone()

	project_deleteCmd.Flags().String("domain", "", "Domain owning <project> (name or ID)")
	projectCmd.AddCommand(project_deleteCmd)
}
