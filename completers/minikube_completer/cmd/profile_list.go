package cmd

import (
	"github.com/spf13/cobra"
)

var profile_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all minikube profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	profile_listCmd.Flags().BoolP("light", "l", false, "If true, returns list of profiles faster by skipping validating the status of the cluster.")
	profile_listCmd.Flags().StringP("output", "o", "table", "The output format. One of 'json', 'table'")
	profileCmd.AddCommand(profile_listCmd)
}
