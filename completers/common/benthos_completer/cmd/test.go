package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Execute Benthos unit tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().String("log", "", "allow components to write logs at a provided level to stdout.")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"log": carapace.ActionValues("error", "warn", "info", "debug", "trace").StyleF(style.ForLogLevel),
	})

	carapace.Gen(testCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
