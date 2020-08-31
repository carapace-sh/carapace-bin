package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	restartCmd.Flags().StringP("timeout", "t", "", "Specify a shutdown timeout in seconds.")
	rootCmd.AddCommand(restartCmd)

	carapace.Gen(restartCmd).PositionalAnyCompletion(ActionServices())
}
