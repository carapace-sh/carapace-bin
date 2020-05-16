package cmd

import (
	"github.com/spf13/cobra"
)

var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage set of tracked repositories",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	remoteCmd.Flags().BoolP("verbose", "v", false, "be verbose; must be placed before a subcommand")
	rootCmd.AddCommand(remoteCmd)
}
