package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pane_reportMetadataCmd = &cobra.Command{
	Use:   "report-metadata <pane_id>",
	Short: "Report display-only pane metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_reportMetadataCmd).Standalone()

	pane_reportMetadataCmd.Flags().String("agent", "", "")
	pane_reportMetadataCmd.Flags().String("applies-to-source", "", "")
	pane_reportMetadataCmd.Flags().Bool("clear-display-agent", false, "")
	pane_reportMetadataCmd.Flags().Bool("clear-state-labels", false, "")
	pane_reportMetadataCmd.Flags().Bool("clear-title", false, "")
	pane_reportMetadataCmd.Flags().StringArray("clear-token", nil, "")
	pane_reportMetadataCmd.Flags().String("display-agent", "", "")
	pane_reportMetadataCmd.Flags().String("seq", "", "")
	pane_reportMetadataCmd.Flags().String("source", "", "")
	pane_reportMetadataCmd.Flags().String("state-label", "", "")
	pane_reportMetadataCmd.Flags().String("title", "", "")
	pane_reportMetadataCmd.Flags().StringArray("token", nil, "")
	pane_reportMetadataCmd.Flags().String("ttl-ms", "", "")
	paneCmd.AddCommand(pane_reportMetadataCmd)

	carapace.Gen(pane_reportMetadataCmd).PositionalCompletion(action.ActionPanes(pane_reportMetadataCmd))
}
