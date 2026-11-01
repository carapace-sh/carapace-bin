package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop all machines in the deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopCmd).Standalone()

	stopCmd.Flags().StringSlice("exclude", nil, "Stop all except the specified machines")
	stopCmd.Flags().StringSlice("include", nil, "Stop only the specified machines")

	rootCmd.AddCommand(stopCmd)
}
