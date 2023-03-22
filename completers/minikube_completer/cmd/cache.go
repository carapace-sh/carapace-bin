package cmd

import (
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:     "cache",
	Short:   "Add, delete, or push a local image into minikube",
	GroupID: "images",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(cacheCmd)
}
