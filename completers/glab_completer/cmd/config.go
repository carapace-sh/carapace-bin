package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set and get glab settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	configCmd.Flags().BoolP("global", "g", false, "use global config file")
	rootCmd.AddCommand(configCmd)
}
