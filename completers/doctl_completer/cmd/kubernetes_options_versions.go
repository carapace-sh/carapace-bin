package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_options_versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List Kubernetes versions that can be used with DigitalOcean clusters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_options_versionsCmd).Standalone()
	kubernetes_optionsCmd.AddCommand(kubernetes_options_versionsCmd)
}
