package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var localsnapshotCmd = &cobra.Command{
	Use:   "localsnapshot",
	Short: "create new local Time Machine snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(localsnapshotCmd).Standalone()
	rootCmd.AddCommand(localsnapshotCmd)
}