package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_user_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve details about a database user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_user_getCmd).Standalone()
	databases_user_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Name`, `Role`, `Password`")
	databases_user_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	databases_userCmd.AddCommand(databases_user_getCmd)
}
