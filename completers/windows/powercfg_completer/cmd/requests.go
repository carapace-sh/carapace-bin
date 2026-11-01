package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var requestsCmd = &cobra.Command{
	Use:   "requests",
	Short: "enumerate application and driver power requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(requestsCmd).Standalone()
	rootCmd.AddCommand(requestsCmd)
}
