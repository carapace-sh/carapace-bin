package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addGenericPasswordCmd = &cobra.Command{
	Use:   "add-generic-password",
	Short: "Add a generic password item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addGenericPasswordCmd).Standalone()
	addGenericPasswordCmd.Flags().StringP("account", "a", "", "Specify account name (required)")
	addGenericPasswordCmd.Flags().BoolP("allow-all", "A", false, "Allow any application to access without warning")
	addGenericPasswordCmd.Flags().StringP("attribute", "G", "", "Specify generic attribute value")
	addGenericPasswordCmd.Flags().StringP("comment", "j", "", "Specify comment string")
	addGenericPasswordCmd.Flags().StringP("creator", "c", "", "Specify item creator (four-character code)")
	addGenericPasswordCmd.Flags().StringP("kind", "D", "", "Specify kind (default: application password)")
	addGenericPasswordCmd.Flags().StringP("label", "l", "", "Specify label (default: service name)")
	addGenericPasswordCmd.Flags().StringP("password", "w", "", "Specify password to be added")
	addGenericPasswordCmd.Flags().StringP("service", "s", "", "Specify service name (required)")
	addGenericPasswordCmd.Flags().StringP("tool", "T", "", "Specify an application which may access this item")
	addGenericPasswordCmd.Flags().StringP("type", "C", "", "Specify item type (four-character code)")
	addGenericPasswordCmd.Flags().BoolP("update", "U", false, "Update item if it already exists")
	rootCmd.AddCommand(addGenericPasswordCmd)
}
