package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_dropletAction_resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize a Droplet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_dropletAction_resizeCmd).Standalone()
	compute_dropletAction_resizeCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_dropletAction_resizeCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_dropletAction_resizeCmd.Flags().Bool("resize-disk", false, "Resize the Droplet's disk size in addition to its RAM and CPU.")
	compute_dropletAction_resizeCmd.Flags().String("size", "", "A slug indicating the new size for the Droplet (e.g. `s-2vcpu-2gb`). Run `doctl compute size list` for a list of valid sizes. (required)")
	compute_dropletAction_resizeCmd.Flags().Bool("wait", false, "Wait for action to complete")
	compute_dropletActionCmd.AddCommand(compute_dropletAction_resizeCmd)

	// TODO size completion
	carapace.Gen(compute_dropletAction_resizeCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
