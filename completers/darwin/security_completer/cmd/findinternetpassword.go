package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var findInternetPasswordCmd = &cobra.Command{
	Use:   "find-internet-password",
	Short: "Find an internet password item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findInternetPasswordCmd).Standalone()
	findInternetPasswordCmd.Flags().StringP("account", "a", "", "Match account string")
	findInternetPasswordCmd.Flags().StringP("auth-type", "t", "", "Match authentication type")
	findInternetPasswordCmd.Flags().StringP("comment", "j", "", "Match comment string")
	findInternetPasswordCmd.Flags().StringP("creator", "c", "", "Match creator (four-character code)")
	findInternetPasswordCmd.Flags().StringP("domain", "d", "", "Match security domain string")
	findInternetPasswordCmd.Flags().BoolP("get-password", "g", false, "Display the password for the item found")
	findInternetPasswordCmd.Flags().StringP("kind", "D", "", "Match kind string")
	findInternetPasswordCmd.Flags().StringP("label", "l", "", "Match label string")
	findInternetPasswordCmd.Flags().BoolP("password-only", "w", false, "Display the password only for the item found")
	findInternetPasswordCmd.Flags().StringP("path", "p", "", "Match path string")
	findInternetPasswordCmd.Flags().StringP("port", "P", "", "Match port number")
	findInternetPasswordCmd.Flags().StringP("protocol", "r", "", "Match protocol (four-character code)")
	findInternetPasswordCmd.Flags().StringP("server", "s", "", "Match server string")
	findInternetPasswordCmd.Flags().StringP("type", "C", "", "Match type (four-character code)")
	rootCmd.AddCommand(findInternetPasswordCmd)
}
