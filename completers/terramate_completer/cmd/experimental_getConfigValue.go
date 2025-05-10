package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var experimental_getConfigValueCmd = &cobra.Command{
	Use:   "get-config-value",
	Short: "Get configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_getConfigValueCmd).Standalone()

	experimental_getConfigValueCmd.Flags().Bool("as-json", false, "Outputs the result as a JSON value")
	experimental_getConfigValueCmd.Flags().StringSliceP("global", "g", nil, "set/override globals. eg.: --global name=<expr>")
	experimentalCmd.AddCommand(experimental_getConfigValueCmd)
}
