package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addons_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all available minikube addons as well as their current statuses (enabled/disabled)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addons_listCmd).Standalone()

	addons_listCmd.Flags().BoolP("docs", "d", false, "If true, print web links to addons' documentation if using --output=list (default).")
	addons_listCmd.Flags().StringP("output", "o", "", "minikube addons list --output OUTPUT. json, list")
	addonsCmd.AddCommand(addons_listCmd)

	carapace.Gen(addons_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("json", "list"),
	})
}
