package cmd

import (
	"github.com/spf13/cobra"
)

var policy_validateConfigCmd = &cobra.Command{
	Use:   "validate-config",
	Short: "Validate a Policy Pack configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	policy_validateConfigCmd.Flags().String("config", "", "The file path for the Policy Pack configuration file")
	policyCmd.AddCommand(policy_validateConfigCmd)
}
