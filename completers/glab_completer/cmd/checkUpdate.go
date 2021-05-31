package cmd

import (
	"github.com/spf13/cobra"
)

var checkUpdateCmd = &cobra.Command{
	Use:   "check-update",
	Short: "Check for latest glab releases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(checkUpdateCmd)
}
