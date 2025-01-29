package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "Play containers, pods or volumes from a structured file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubeCmd).Standalone()

	rootCmd.AddCommand(kubeCmd)
}
