package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen for notifications",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listenCmd).Standalone()
	listenCmd.Flags().String("addasn", "", "Add ASN")
	listenCmd.Flags().String("id", "", "Notification ID")
	listenCmd.Flags().String("removeasn", "", "Remove ASN")
	rootCmd.AddCommand(listenCmd)
}
