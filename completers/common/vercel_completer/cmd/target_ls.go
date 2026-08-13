package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var target_lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List targets defined for the current Project",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(target_lsCmd).Standalone()

	target_lsCmd.Flags().String("format", "", "Output format")
	target_lsCmd.Flags().Bool("json", false, "Output as JSON")
	target_lsCmd.Flags().String("project", "", "Project name or ID")
	target_lsCmd.Flags().Bool("yes", false, "Skip confirmation")

	targetCmd.AddCommand(target_lsCmd)
}
