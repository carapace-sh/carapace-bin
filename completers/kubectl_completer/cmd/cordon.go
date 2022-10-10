package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cordonCmd = &cobra.Command{
	Use:   "cordon",
	Short: "Mark node as unschedulable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cordonCmd).Standalone()
	cordonCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	cordonCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
	cordonCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	rootCmd.AddCommand(cordonCmd)

	carapace.Gen(cordonCmd).FlagCompletion(carapace.ActionMap{
		"dry-run": action.ActionDryRunModes(),
	})

	carapace.Gen(cordonCmd).PositionalCompletion(
		action.ActionResources("", "nodes"),
	)
}