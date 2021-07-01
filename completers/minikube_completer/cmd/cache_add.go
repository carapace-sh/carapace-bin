package cmd

import (
	"github.com/spf13/cobra"
)

var cache_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an image to local cache.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	cache_addCmd.Flags().Bool("", false, "Add image to cache for all running minikube clusters")
	cacheCmd.AddCommand(cache_addCmd)
}
