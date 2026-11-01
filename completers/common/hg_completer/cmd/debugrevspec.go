package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugrevspecCmd = &cobra.Command{
	Use:    "debugrevspec",
	Short:  "REVSPEC",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugrevspecCmd).Standalone()

	debugrevspecCmd.Flags().Bool("no-optimized", false, "evaluate tree without optimization")
	debugrevspecCmd.Flags().Bool("optimize", false, "print parsed tree after optimizing (DEPRECATED)")
	debugrevspecCmd.Flags().Bool("show-revs", false, "print list of result revisions (default)")
	debugrevspecCmd.Flags().BoolP("show-set", "s", false, "print internal representation of result set")
	debugrevspecCmd.Flags().StringArrayP("show-stage", "p", nil, "print parsed tree at the given stage")
	debugrevspecCmd.Flags().Bool("verify-optimized", false, "verify optimized result")
	rootCmd.AddCommand(debugrevspecCmd)
}
