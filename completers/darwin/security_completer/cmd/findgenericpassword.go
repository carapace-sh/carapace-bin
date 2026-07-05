package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var findGenericPasswordCmd = &cobra.Command{
	Use:   "find-generic-password",
	Short: "Find a generic password item",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findGenericPasswordCmd).Standalone()
	findGenericPasswordCmd.Flags().StringP("account", "a", "", "Match account string")
	findGenericPasswordCmd.Flags().StringP("attribute", "G", "", "Match value string")
	findGenericPasswordCmd.Flags().StringP("comment", "j", "", "Match comment string")
	findGenericPasswordCmd.Flags().StringP("creator", "c", "", "Match creator (four-character code)")
	findGenericPasswordCmd.Flags().BoolP("get-password", "g", false, "Display the password for the item found")
	findGenericPasswordCmd.Flags().StringP("kind", "D", "", "Match kind string")
	findGenericPasswordCmd.Flags().StringP("label", "l", "", "Match label string")
	findGenericPasswordCmd.Flags().BoolP("password-only", "w", false, "Display the password only for the item found")
	findGenericPasswordCmd.Flags().StringP("service", "s", "", "Match service string")
	findGenericPasswordCmd.Flags().StringP("type", "C", "", "Match type (four-character code)")
	rootCmd.AddCommand(findGenericPasswordCmd)
}
