package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().String("args", "", "How many arguments to expect in the exec command")
	execCmd.Flags().Bool("sync", false, "Block until the child process exits")
	execCmd.Flags().String("terminator", "", "Specifies a terminator that marks the end of exec arguments")
	rootCmd.AddCommand(execCmd)
}
