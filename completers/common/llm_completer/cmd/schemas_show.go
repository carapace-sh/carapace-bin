package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a stored schema",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_showCmd).Standalone()

	schemas_showCmd.Flags().StringP("database", "d", "", "Path to log database")
	schemas_showCmd.Flags().StringP("path", "p", "", "Path to log database")
	schemasCmd.AddCommand(schemas_showCmd)
}
