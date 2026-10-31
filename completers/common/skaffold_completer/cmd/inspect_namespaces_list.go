package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspect_namespaces_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the list of namespaces that would be run for a given configuration (default skaffold configuration, specific module, specific profile, etc).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspect_namespaces_listCmd).Standalone()

	inspect_namespaces_listCmd.Flags().StringSliceP("module", "m", nil, "Names of modules to filter target action by.")
	inspect_namespaces_listCmd.Flags().StringSliceP("profile", "p", nil, "Profile names to activate")
	inspect_namespaces_listCmd.Flags().Bool("propagate-profiles", true, "Setting '--propagate-profiles=false' disables propagating profiles set by the '--profile' flag across config dependencies. This mean that only profiles defined directly in the target 'skaffold.yaml' file are activated.")
	applyCommonFlagCompletion(inspect_namespaces_listCmd)
	inspect_namespacesCmd.AddCommand(inspect_namespaces_listCmd)
}
