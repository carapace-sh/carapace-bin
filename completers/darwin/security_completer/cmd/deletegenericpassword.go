package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteGenericPasswordCmd = &cobra.Command{
	Use:   "delete-generic-password",
	Short: "Delete a generic password item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteGenericPasswordCmd).Standalone()
	deleteGenericPasswordCmd.Flags().StringP("account", "a", "", "Match account string")
	deleteGenericPasswordCmd.Flags().StringP("attribute", "G", "", "Match value string")
	deleteGenericPasswordCmd.Flags().StringP("comment", "j", "", "Match comment string")
	deleteGenericPasswordCmd.Flags().StringP("creator", "c", "", "Match creator (four-character code)")
	deleteGenericPasswordCmd.Flags().StringP("kind", "D", "", "Match kind string")
	deleteGenericPasswordCmd.Flags().StringP("label", "l", "", "Match label string")
	deleteGenericPasswordCmd.Flags().StringP("service", "s", "", "Match service string")
	deleteGenericPasswordCmd.Flags().StringP("type", "C", "", "Match type (four-character code)")
	rootCmd.AddCommand(deleteGenericPasswordCmd)
}
