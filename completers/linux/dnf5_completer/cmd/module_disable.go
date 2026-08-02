package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var moduleDisableCmd = &cobra.Command{
	Use:   "disable [options] <module-spec>...",
	Short: "disable modules including all their streams",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleDisableCmd).Standalone()

	moduleDisableCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")

	moduleCmd.AddCommand(moduleDisableCmd)
}
