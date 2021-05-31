package cmd

import (
	"github.com/spf13/cobra"
)

var config_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Shows a prompt to set basic glab configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	configCmd.AddCommand(config_initCmd)
}
