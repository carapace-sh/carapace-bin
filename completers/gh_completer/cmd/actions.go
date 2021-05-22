package cmd

import (
	"github.com/spf13/cobra"
)

var actionsCmd = &cobra.Command{
	Use:   "actions",
	Short: "Learn about working with GitHub actions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(actionsCmd)
}
