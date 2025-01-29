package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_existsCmd = &cobra.Command{
	Use:   "exists POD",
	Short: "Check if a pod exists in local storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_existsCmd).Standalone()

	podCmd.AddCommand(pod_existsCmd)
}
