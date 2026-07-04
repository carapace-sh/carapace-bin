package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var leaksCmd = &cobra.Command{
	Use:   "leaks",
	Short: "Run leaks on this process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(leaksCmd).Standalone()
	leaksCmd.Flags().Bool("cycles", false, "Use a stricter algorithm")
	leaksCmd.Flags().String("exclude", "", "Ignore leaks called from symbol")
	leaksCmd.Flags().Bool("nocontext", false, "Withhold hex dumps of leaked memory")
	leaksCmd.Flags().Bool("nostacks", false, "Do not show stack traces of leaked memory")
	rootCmd.AddCommand(leaksCmd)
}
