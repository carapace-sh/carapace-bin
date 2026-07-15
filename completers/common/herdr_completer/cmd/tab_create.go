package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tab_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tab_createCmd).Standalone()

	tab_createCmd.Flags().String("cwd", "", "")
	tab_createCmd.Flags().StringSlice("env", nil, "Set an environment variable for the launched process")
	tab_createCmd.Flags().Bool("focus", false, "")
	tab_createCmd.Flags().String("label", "", "")
	tab_createCmd.Flags().Bool("no-focus", false, "")
	tab_createCmd.Flags().String("workspace", "", "")
	tabCmd.AddCommand(tab_createCmd)

	carapace.Gen(tab_createCmd).FlagCompletion(carapace.ActionMap{
		"cwd": carapace.ActionFiles(),
	})
}
