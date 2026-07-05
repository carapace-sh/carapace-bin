package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addInternetPasswordCmd = &cobra.Command{
	Use:   "add-internet-password",
	Short: "Add an internet password item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addInternetPasswordCmd).Standalone()
	addInternetPasswordCmd.Flags().StringP("account", "a", "", "Specify account name (required)")
	addInternetPasswordCmd.Flags().BoolP("allow-all", "A", false, "Allow any application to access without warning")
	addInternetPasswordCmd.Flags().StringP("auth-type", "t", "", "Specify authentication type (default: dflt)")
	addInternetPasswordCmd.Flags().StringP("comment", "j", "", "Specify comment string")
	addInternetPasswordCmd.Flags().StringP("creator", "c", "", "Specify item creator (four-character code)")
	addInternetPasswordCmd.Flags().StringP("domain", "d", "", "Specify security domain string")
	addInternetPasswordCmd.Flags().StringP("kind", "D", "", "Specify kind (default: application password)")
	addInternetPasswordCmd.Flags().StringP("label", "l", "", "Specify label (default: service name)")
	addInternetPasswordCmd.Flags().StringP("password", "w", "", "Specify password to be added")
	addInternetPasswordCmd.Flags().StringP("path", "p", "", "Specify path string")
	addInternetPasswordCmd.Flags().StringP("port", "P", "", "Specify port number")
	addInternetPasswordCmd.Flags().StringP("protocol", "r", "", "Specify protocol (four-character code)")
	addInternetPasswordCmd.Flags().StringP("server", "s", "", "Specify server name (required)")
	addInternetPasswordCmd.Flags().StringP("tool", "T", "", "Specify an application which may access this item")
	addInternetPasswordCmd.Flags().StringP("type", "C", "", "Specify item type (four-character code)")
	addInternetPasswordCmd.Flags().BoolP("update", "U", false, "Update item if it already exists")
	rootCmd.AddCommand(addInternetPasswordCmd)
}
