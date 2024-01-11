package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var waitCmd = &cobra.Command{
	Use:   "wait SERVICE [SERVICE...] [OPTIONS]",
	Short: "Block until the first service container stops",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitCmd).Standalone()

	waitCmd.Flags().Bool("down-project", false, "Drops project when the first container stops")
	rootCmd.AddCommand(waitCmd)
}
