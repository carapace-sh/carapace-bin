package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:     "configure",
	Aliases: []string{"configuration", "dsc"},
	Short:   "Configures the system into a desired state",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configureCmd).Standalone()

	configureCmd.PersistentFlags().StringP("file", "f", "", "The path to the configuration file")
	configureCmd.PersistentFlags().String("module-path", "", "Specifies the location on the local computer to store modules")
	configureCmd.PersistentFlags().String("processor-path", "", "Specify the path to the configuration processor")
	rootCmd.AddCommand(configureCmd)

	carapace.Gen(configureCmd).FlagCompletion(carapace.ActionMap{
		"file":           carapace.ActionFiles(),
		"module-path":    carapace.ActionFiles(),
		"processor-path": carapace.ActionFiles(),
	})
}
