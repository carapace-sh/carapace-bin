package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "auditpol",
	Short: "display and manipulate audit policies",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/auditpol",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"get", "display the current audit policy",
			"set", "set the audit policy",
			"list", "list audit policy categories",
			"backup", "backup the audit policy",
			"restore", "restore the audit policy",
			"clear", "clear the audit policy",
			"remove", "remove a per-user audit policy",
		),
	)
}
