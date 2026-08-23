package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(enumerateCmd)
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(materializeCmd)
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(dumpCmd)
}

var enumerateCmd = &cobra.Command{
	Use:   "enumerate",
	Short: "run an interactive enumeration of the specified provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "run an interactive enumeration of the specified provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var materializeCmd = &cobra.Command{
	Use:   "materialize",
	Short: "cause the specified item to be written on disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "run the validation suite against the specified provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "dump the state of the file provider subsystem",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enumerateCmd).Standalone()
	carapace.Gen(lsCmd).Standalone()
	carapace.Gen(materializeCmd).Standalone()
	carapace.Gen(validateCmd).Standalone()
	carapace.Gen(dumpCmd).Standalone()
}