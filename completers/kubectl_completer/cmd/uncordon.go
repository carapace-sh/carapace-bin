package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var uncordonCmd = &cobra.Command{
	Use:     "uncordon NODE",
	Short:   "Mark node as schedulable",
	GroupID: "cluster management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uncordonCmd).Standalone()

	uncordonCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	uncordonCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
	uncordonCmd.Flag("dry-run").NoOptDefVal = " "
	rootCmd.AddCommand(uncordonCmd)

	carapace.Gen(uncordonCmd).FlagCompletion(carapace.ActionMap{
		"dry-run": kubectl.ActionDryRunModes(),
	})

	carapace.Gen(uncordonCmd).PositionalCompletion(
		kubectl.ActionResources(kubectl.ResourceOpts{Namespace: "", Types: "nodes"}),
	)
}
