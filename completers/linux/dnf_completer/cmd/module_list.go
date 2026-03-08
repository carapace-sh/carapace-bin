package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var moduleListCmd = &cobra.Command{
	Use:   "list [module-spec]...",
	Short: "list module streams",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleListCmd).Standalone()

	moduleListCmd.Flags().Bool("disabled", false, "show only disabled modules")
	moduleListCmd.Flags().Bool("enabled", false, "show only enabled modules")

	moduleCmd.AddCommand(moduleListCmd)
}
