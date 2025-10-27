package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var profile_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all minikube profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(profile_listCmd).Standalone()

	profile_listCmd.Flags().BoolP("detailed", "d", false, "If true, returns a detailed list of profiles.")
	profile_listCmd.Flags().BoolP("light", "l", false, "If true, returns list of profiles faster by skipping validating the status of the cluster.")
	profile_listCmd.Flags().StringP("output", "o", "", "The output format. One of 'json', 'table'")
	profileCmd.AddCommand(profile_listCmd)

	carapace.Gen(profile_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("json", "table"),
	})
}
