package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var azCmd = &cobra.Command{
	Use:   "az",
	Short: "Access Azure API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(azCmd).Standalone()

	azCmd.Flags().String("app", "", "Optional name of the Azure application to use if logged into multiple.")
	azCmd.Flags().String("azure-identity", "", "(For Azure CLI access only) Azure managed identity name.")
	rootCmd.AddCommand(azCmd)
}
