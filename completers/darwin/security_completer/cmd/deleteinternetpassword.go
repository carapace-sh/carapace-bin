package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteInternetPasswordCmd = &cobra.Command{
	Use:   "delete-internet-password",
	Short: "Delete an internet password item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteInternetPasswordCmd).Standalone()
	deleteInternetPasswordCmd.Flags().StringP("account", "a", "", "Match account string")
	deleteInternetPasswordCmd.Flags().StringP("auth-type", "t", "", "Match authentication type")
	deleteInternetPasswordCmd.Flags().StringP("comment", "j", "", "Match comment string")
	deleteInternetPasswordCmd.Flags().StringP("creator", "c", "", "Match creator (four-character code)")
	deleteInternetPasswordCmd.Flags().StringP("domain", "d", "", "Match security domain string")
	deleteInternetPasswordCmd.Flags().StringP("kind", "D", "", "Match kind string")
	deleteInternetPasswordCmd.Flags().StringP("label", "l", "", "Match label string")
	deleteInternetPasswordCmd.Flags().StringP("path", "p", "", "Match path string")
	deleteInternetPasswordCmd.Flags().StringP("port", "P", "", "Match port number")
	deleteInternetPasswordCmd.Flags().StringP("protocol", "r", "", "Match protocol (four-character code)")
	deleteInternetPasswordCmd.Flags().StringP("server", "s", "", "Match server string")
	deleteInternetPasswordCmd.Flags().StringP("type", "C", "", "Match type (four-character code)")
	rootCmd.AddCommand(deleteInternetPasswordCmd)
}
