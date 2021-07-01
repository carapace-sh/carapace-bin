package cmd

import (
	"github.com/spf13/cobra"
)

var service_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the URLs for the services in your local cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	service_listCmd.Flags().StringP("namespace", "n", "", "The services namespace")
	serviceCmd.AddCommand(service_listCmd)
}
