package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runOperationCmd = &cobra.Command{
	Use:   "run-operation",
	Short: "Run the named macro with any supplied arguments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runOperationCmd).Standalone()

	runOperationCmd.Flags().String("args", "", "Supply arguments to the macro")
	runOperationCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	runOperationCmd.Flags().String("vars", "", "Supply variables to the project")
	addProjectDirFlag(runOperationCmd)
	addProfileFlag(runOperationCmd)
	rootCmd.AddCommand(runOperationCmd)

	// TODO macro
}
