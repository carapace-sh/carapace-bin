package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_sqlMode_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a MySQL database cluster's SQL modes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_sqlMode_getCmd).Standalone()
	databases_sqlMode_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Name`")
	databases_sqlMode_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databases_sqlModeCmd.AddCommand(databases_sqlMode_getCmd)
}
