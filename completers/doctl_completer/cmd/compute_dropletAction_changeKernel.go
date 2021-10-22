package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_dropletAction_changeKernelCmd = &cobra.Command{
	Use:   "change-kernel",
	Short: "Change a Droplet's kernel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_dropletAction_changeKernelCmd).Standalone()
	compute_dropletAction_changeKernelCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Status`, `Type`, `StartedAt`, `CompletedAt`, `ResourceID`, `ResourceType`, `Region`")
	compute_dropletAction_changeKernelCmd.Flags().Int("kernel-id", 0, "Kernel ID (required)")
	compute_dropletAction_changeKernelCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_dropletAction_changeKernelCmd.Flags().Bool("wait", false, "Wait for action to complete")
	compute_dropletActionCmd.AddCommand(compute_dropletAction_changeKernelCmd)

	carapace.Gen(compute_dropletAction_changeKernelCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Status", "Type", "StartedAt", "CompletedAt", "ResourceID", "ResourceType", "Region").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
