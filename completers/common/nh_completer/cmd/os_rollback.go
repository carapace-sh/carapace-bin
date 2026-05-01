package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var os_rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback to a previous generation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(os_rollbackCmd).Standalone()

	os_rollbackCmd.Flags().BoolP("ask", "a", false, "Ask for confirmation")
	os_rollbackCmd.Flags().BoolP("bypass-root-check", "R", false, "Don't panic if calling nh as root")
	os_rollbackCmd.Flags().StringP("diff", "d", "auto", "Whether to display a package diff")
	os_rollbackCmd.Flags().BoolP("dry", "n", false, "Only print actions, without performing them")
	os_rollbackCmd.Flags().BoolP("no-specialisation", "S", false, "Ignore specialisations")
	os_rollbackCmd.Flags().StringP("specialisation", "s", "", "Explicitly select some specialisation")
	os_rollbackCmd.Flags().StringP("to", "t", "", "Rollback to a specific generation number")

	carapace.Gen(os_rollbackCmd).FlagCompletion(carapace.ActionMap{
		"diff": carapace.ActionValues("auto", "always", "never"),
	})

	osCmd.AddCommand(os_rollbackCmd)
}
