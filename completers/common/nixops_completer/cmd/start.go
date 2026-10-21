package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start all machines in the deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()

	startCmd.Flags().StringSlice("exclude", nil, "Start all except the specified machines")
	startCmd.Flags().StringSlice("include", nil, "Start only the specified machines")

	rootCmd.AddCommand(startCmd)
}
