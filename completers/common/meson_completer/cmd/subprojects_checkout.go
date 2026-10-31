package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subprojects_checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout a branch (git only)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subprojects_checkoutCmd).Standalone()

	subprojects_checkoutCmd.Flags().Bool("allow-insecure", false, "Allow insecure server connections.")
	subprojects_checkoutCmd.Flags().BoolS("b", "b", false, "Create a new branch")
	subprojects_checkoutCmd.Flags().StringP("num-processes", "j", "", "How many parallel processes to use (Since 0.59.0).")
	subprojects_checkoutCmd.Flags().String("sourcedir", "", "Path to source directory")
	subprojects_checkoutCmd.Flags().String("types", "", "Comma-separated list of subproject types.")
	applyCommonFlagCompletion(subprojects_checkoutCmd)
	subprojectsCmd.AddCommand(subprojects_checkoutCmd)
}
