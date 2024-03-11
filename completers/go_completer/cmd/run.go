package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "compile and run Go program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().SetInterspersed(false)

	runCmd.Flags().StringS("exec", "exec", "", "invoke the binary using xprog")
	addBuildFlags(runCmd)
	rootCmd.AddCommand(runCmd)
}
