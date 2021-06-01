package cmd

import (
	"github.com/spf13/cobra"
)

var policy_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a Policy Pack to the Pulumi service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	policyCmd.AddCommand(policy_publishCmd)
}
