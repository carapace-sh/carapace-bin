package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compute_imageAction_transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer an image to another datacenter region",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_imageAction_transferCmd).Standalone()
	compute_imageAction_transferCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_imageAction_transferCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_imageAction_transferCmd.Flags().String("region", "", "region (required)")
	compute_imageAction_transferCmd.Flags().Bool("wait", false, "Wait for action to complete")
	compute_imageActionCmd.AddCommand(compute_imageAction_transferCmd)

	carapace.Gen(compute_imageAction_transferCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
		"region": action.ActionRegions(),
	})
}
