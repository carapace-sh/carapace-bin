package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var databases_sqlMode_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a MySQL database cluster's SQL modes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databases_sqlMode_setCmd).Standalone()
	databases_sqlModeCmd.AddCommand(databases_sqlMode_setCmd)
}
