package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Serve a source over the network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverCmd).Standalone()
	serverCmd.Flags().String("source", "", "Source image to serve")
	rootCmd.AddCommand(serverCmd)
	carapace.Gen(serverCmd).FlagCompletion(carapace.ActionMap{
		"source": carapace.ActionFiles(),
	})
}
