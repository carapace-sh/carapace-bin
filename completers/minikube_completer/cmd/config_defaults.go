package cmd

import (
	"github.com/spf13/cobra"
)

var config_defaultsCmd = &cobra.Command{
	Use:   "defaults",
	Short: "Lists all valid default values for PROPERTY_NAME",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	config_defaultsCmd.Flags().String("output", "table", "Output format. Accepted values: [json]")
	configCmd.AddCommand(config_defaultsCmd)
}
