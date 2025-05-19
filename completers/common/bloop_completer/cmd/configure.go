package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configure projecs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configureCmd).Standalone()

	configureCmd.Flags().String("threads", "", "set the number of threads used to compile and test all projects")
	rootCmd.AddCommand(configureCmd)
}
