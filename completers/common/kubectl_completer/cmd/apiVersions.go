package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apiVersionsCmd = &cobra.Command{
	Use:   "api-versions",
	Short: "Print the supported API versions on the server, in the form of \"group/version\"",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiVersionsCmd).Standalone()

	rootCmd.AddCommand(apiVersionsCmd)
}
