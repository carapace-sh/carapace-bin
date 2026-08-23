package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nettop",
	Short: "display updated information about the network",
	Long:  "https://keith.github.io/xcode-manpages/nettop.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("J", "J", "", "Column selection")
	rootCmd.Flags().StringS("L", "L", "", "Samples")
	rootCmd.Flags().BoolS("P", "P", false, "Delta output")
	rootCmd.Flags().BoolS("c", "c", false, "Delta output")
	rootCmd.Flags().BoolS("d", "d", false, "Delta output")
	rootCmd.Flags().StringS("j", "j", "", "Column selection")
	rootCmd.Flags().StringS("k", "k", "", "Column selection")
	rootCmd.Flags().StringS("l", "l", "", "Samples")
	rootCmd.Flags().StringS("m", "m", "", "Mode")
	rootCmd.Flags().BoolS("n", "n", false, "Numeric output")
	rootCmd.Flags().StringS("p", "p", "", "Process name or PID")
	rootCmd.Flags().StringS("s", "s", "", "Refresh interval")
	rootCmd.Flags().StringS("t", "t", "", "Type")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"m": carapace.ActionValues("tcp", "udp", "scopt", "route"),
		"p": ps.ActionProcessExecutables(),
		"t": carapace.ActionValues("tcp", "udp", "scopt", "route"),
	})
}
