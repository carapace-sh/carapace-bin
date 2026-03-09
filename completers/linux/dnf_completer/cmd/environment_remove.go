package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environment_removeCmd = &cobra.Command{
	Use:   "remove [options] <group-spec|environment-spec>...",
	Short: "remove environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environment_removeCmd)

	environment_removeCmd.Flags().Bool("no-packages", false, "operate on environments without manipulating packages")
	environment_removeCmd.Flags().Bool("with-optional", false, "include optional packages")
	environmentCmd.AddCommand(environment_removeCmd)
}
