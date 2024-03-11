package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var describeStoresCmd = &cobra.Command{
	Use:     "describe-stores",
	Short:   "show registered store types and their available options",
	GroupID: "utility",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(describeStoresCmd).Standalone()

	describeStoresCmd.Flags().Bool("json", false, "Produce output in JSON format")
	rootCmd.AddCommand(describeStoresCmd)

	// TODO positional completion
}
