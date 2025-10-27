package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateContextCmd = &cobra.Command{
	Use:     "update-context",
	Short:   "Update kubeconfig in case of an IP or port change",
	GroupID: "configuration",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateContextCmd).Standalone()

	rootCmd.AddCommand(updateContextCmd)
}
