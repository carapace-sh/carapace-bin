package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config SUBCOMMAND",
	Short: "Modify kubeconfig files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.PersistentFlags().String("kubeconfig", "", "use a particular kubeconfig file")
	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"kubeconfig": carapace.ActionFiles(),
	})
}
