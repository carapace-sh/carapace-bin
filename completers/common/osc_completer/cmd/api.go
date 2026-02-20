package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Perform direct API requests to OpenStack services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiCmd).Standalone()

	apiCmd.Flags().String("service", "", "Service type")
	apiCmd.Flags().String("url", "", "URL path")
	apiCmd.Flags().String("method", "", "HTTP method")
	apiCmd.Flags().String("body", "", "Request body")

	rootCmd.AddCommand(apiCmd)

	carapace.Gen(apiCmd).FlagCompletion(carapace.ActionMap{
		"method": carapace.ActionValues("GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"),
	})
}
