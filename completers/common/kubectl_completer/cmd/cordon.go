package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var cordonCmd = &cobra.Command{
	Use:     "cordon NODE",
	Short:   "Mark node as unschedulable",
	GroupID: "cluster management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cordonCmd).Standalone()

	cordonCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	cordonCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	cordonCmd.Flag("dry-run").NoOptDefVal = " "
	rootCmd.AddCommand(cordonCmd)

	carapace.Gen(cordonCmd).FlagCompletion(carapace.ActionMap{
		"dry-run": kubectl.ActionDryRunModes(),
	})

	carapace.Gen(cordonCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionResources(kubectl.ResourceOpts{
				Context:   rootCmd.Flag("context").Value.String(),
				Namespace: rootCmd.Flag("namespace").Value.String(),
				Types:     "nodes",
			})
		}),
	)
}
