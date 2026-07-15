package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_reportMetadataCmd = &cobra.Command{
	Use:   "report-metadata <workspace_id>",
	Short: "Report display-only workspace metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_reportMetadataCmd).Standalone()

	workspace_reportMetadataCmd.Flags().StringSlice("clear-token", nil, "")
	workspace_reportMetadataCmd.Flags().String("seq", "", "")
	workspace_reportMetadataCmd.Flags().String("source", "", "")
	workspace_reportMetadataCmd.Flags().StringSlice("token", nil, "")
	workspace_reportMetadataCmd.Flags().String("ttl-ms", "", "")
	workspaceCmd.AddCommand(workspace_reportMetadataCmd)
}
