package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:     "cache",
	Short:   "Add, delete, or push a local image into minikube",
	GroupID: "images",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cacheCmd).Standalone()
	rootCmd.AddCommand(cacheCmd)
}
