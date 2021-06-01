package cmd

import (
	"github.com/spf13/cobra"
)

var config_refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Update the local configuration based on the most recent deployment of the stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	config_refreshCmd.PersistentFlags().BoolP("force", "f", false, "Overwrite configuration file, if it exists, without creating a backup")
	configCmd.AddCommand(config_refreshCmd)
}
