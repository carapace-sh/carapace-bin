package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var drainCmd = &cobra.Command{
	Use:     "drain NODE",
	Short:   "Drain node in preparation for maintenance",
	GroupID: "cluster management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drainCmd).Standalone()

	drainCmd.Flags().String("chunk-size", "", "Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future.")
	drainCmd.Flags().Bool("delete-emptydir-data", false, "Continue even if there are pods using emptyDir (local data that will be deleted when the node is drained).")
	drainCmd.Flags().Bool("disable-eviction", false, "Force drain to use delete, even if eviction is supported. This will bypass checking PodDisruptionBudgets, use with caution.")
	drainCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	drainCmd.Flags().Bool("force", false, "Continue even if there are pods that do not declare a controller.")
	drainCmd.Flags().String("grace-period", "", "Period of time in seconds given to each pod to terminate gracefully. If negative, the default value specified in the pod will be used.")
	drainCmd.Flags().Bool("ignore-daemonsets", false, "Ignore DaemonSet-managed pods.")
	drainCmd.Flags().String("pod-selector", "", "Label selector to filter pods on the node")
	drainCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	drainCmd.Flags().String("skip-wait-for-delete-timeout", "", "If pod DeletionTimestamp older than N seconds, skip waiting for the pod.  Seconds must be greater than 0 to skip.")
	drainCmd.Flags().String("timeout", "", "The length of time to wait before giving up, zero means infinite")
	drainCmd.Flag("dry-run").NoOptDefVal = " "
	rootCmd.AddCommand(drainCmd)

	carapace.Gen(drainCmd).FlagCompletion(carapace.ActionMap{
		"dry-run": kubectl.ActionDryRunModes(),
	})

	carapace.Gen(drainCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionResources(kubectl.ResourceOpts{
				Context:   rootCmd.Flag("context").Value.String(),
				Namespace: rootCmd.Flag("namespace").Value.String(),
				Types:     "nodes",
			})
		}),
	)
}
