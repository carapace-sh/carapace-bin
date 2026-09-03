package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var debug_revsetCmd = &cobra.Command{
	Use:   "revset",
	Short: "Evaluate revset to full commit IDs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_revsetCmd).Standalone()

	debug_revsetCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_revsetCmd.Flags().Bool("no-optimize", false, "Do not rewrite expression to optimized form")
	debug_revsetCmd.Flags().Bool("no-resolve", false, "Do not resolve and evaluate expression")
	debugCmd.AddCommand(debug_revsetCmd)

	carapace.Gen(debug_revsetCmd).PositionalCompletion(
		jj.ActionRevsets(jj.RevOpts{}.Default()),
	)
}
