package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_imageAction_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve the status of an image action",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_imageAction_getCmd).Standalone()
	compute_imageAction_getCmd.Flags().Int("action-id", 0, "action id (required)")
	compute_imageAction_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_imageAction_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_imageActionCmd.AddCommand(compute_imageAction_getCmd)

	carapace.Gen(compute_imageAction_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Type", "Distribution", "Slug", "Public", "MinDisk").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
